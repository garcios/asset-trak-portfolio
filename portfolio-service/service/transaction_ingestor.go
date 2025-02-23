package service

import (
	"context"
	"fmt"
	"github.com/Thiht/transactor"
	"github.com/garcios/asset-trak-portfolio/lib/excel"
	"github.com/garcios/asset-trak-portfolio/lib/typesutils"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/db"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/model"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

const (
	fieldLength = 8
)

type allCache struct {
	assets *cache.Cache
}

const (
	defaultExpiration = 5 * time.Minute
	purgeTime         = 10 * time.Minute
)

type ITransactionManager interface {
	AddTransaction(ctx context.Context, rec *model.Transaction) error
	Truncate(ctx context.Context) error
}

type IAccountManager interface {
	FindAccountByID(id string) (*model.Account, error)
}

type IBalanceManager interface {
	AddBalance(ctx context.Context, rec *model.AssetBalance) error
	UpdateBalance(ctx context.Context, rec *model.AssetBalance) error
	GetBalance(ctx context.Context, accountID string, assetID string) (*model.AssetBalance, error)
	Truncate(ctx context.Context) error
}

type IAssetManager interface {
	FindAssetBySymbol(symbol string) (*model.Asset, error)
}

// verify interface compliance
var _ ITransactionManager = &db.TransactionRepository{}
var _ IAccountManager = &db.AccountRepository{}
var _ IBalanceManager = &db.AssetBalanceRepository{}

type TransactionIngestor struct {
	TransactionManager ITransactionManager
	AccountManager     IAccountManager
	AssetManager       IAssetManager
	BalanceManager     IBalanceManager
	Transactor         transactor.Transactor
	cache              *allCache
}

func NewTransactionIngestor(
	tm ITransactionManager,
	am IAccountManager,
	astm IAssetManager,
	bm IBalanceManager,
	tr transactor.Transactor,
) *TransactionIngestor {
	ac := cache.New(defaultExpiration, purgeTime)
	return &TransactionIngestor{
		TransactionManager: tm,
		AccountManager:     am,
		AssetManager:       astm,
		BalanceManager:     bm,
		Transactor:         tr,
		cache:              &allCache{assets: ac},
	}
}

func (ingestor *TransactionIngestor) Truncate(ctx context.Context) error {
	err := ingestor.BalanceManager.Truncate(ctx)
	if err != nil {
		return fmt.Errorf("failed to truncate balance data: %w", err)
	}

	log.Println("truncated balance data successfully")

	err = ingestor.TransactionManager.Truncate(ctx)
	if err != nil {
		return fmt.Errorf("failed to truncate transaction data: %w", err)
	}

	log.Println("truncated transaction data successfully")

	return nil
}

func (ingestor *TransactionIngestor) ProcessTransactions(
	ctx context.Context,
	filePath string,
	tabName string,
	skipRows int,
	accountID string,
) error {
	log.Println("Processing transactions...")

	// verify account
	account, err := ingestor.AccountManager.FindAccountByID(accountID)
	if err != nil {
		return err
	}

	if account == nil {
		return fmt.Errorf("account with ID %s not found", accountID)
	}

	rows, err := excel.GetRows(filePath, tabName)
	if err != nil {
		return err
	}

	var rowCount int
	for _, row := range rows {
		if rowCount < skipRows {
			rowCount++
			continue
		}

		if len(row) == 0 || row[0] == "Total" {
			break
		}

		transaction, err := ingestor.mapColumnsToTransaction(row)
		if err != nil {
			return err
		}

		// populate IDs
		transaction.ID = uuid.New().String()
		transaction.AccountID = accountID

		err = ingestor.addTransaction(ctx, transaction)
		if err != nil {
			return err
		}

	}

	return nil
}

func (ingestor *TransactionIngestor) mapColumnsToTransaction(row []string) (*model.Transaction, error) {
	if len(row) < fieldLength {
		displayRow(row)
		return nil, fmt.Errorf("row contains less than expected fields: %d", len(row))
	}

	asset, err := ingestor.getAsset(row[0])
	if err != nil {
		displayRow(row)
		return nil, fmt.Errorf("unable to retrieve asset with symbol: %s", row[0])
	}

	transactionDate, err := typesutils.GetDateValue(row[3])
	if err != nil {
		displayRow(row)
		return nil, fmt.Errorf("unable to process transaction date with value: %s", row[3])
	}

	quantity, err := typesutils.GetFloatValue(row[5])
	if err != nil {
		displayRow(row)
		return nil, fmt.Errorf("unable to process quantity with value: %s", row[5])
	}

	price, err := typesutils.GetFloatValue(row[6])
	if err != nil {
		displayRow(row)
		return nil, fmt.Errorf("unable to process price with value: %s", row[6])
	}

	return &model.Transaction{
		AssetID:                asset.ID,
		TransactionType:        typesutils.GetStringValue(row[4]),
		TransactionDate:        transactionDate,
		Quantity:               quantity,
		TradePrice:             price,
		AssetPriceCurrencyCode: row[7],
	}, nil
}

func (ingestor *TransactionIngestor) addTransaction(ctx context.Context, rec *model.Transaction) error {
	// execute add transaction and add/update balance within a transaction.
	return ingestor.Transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		err := ingestor.TransactionManager.AddTransaction(ctx, rec)
		if err != nil {
			return err
		}

		balance, err := ingestor.BalanceManager.GetBalance(ctx, rec.AccountID, rec.AssetID)
		if err != nil {
			return err
		}

		if balance == nil {
			err = ingestor.BalanceManager.AddBalance(ctx, &model.AssetBalance{
				AccountID: rec.AccountID,
				AssetID:   rec.AssetID,
				Quantity:  rec.Quantity,
			})
			if err != nil {
				return err
			}
		} else {
			err = ingestor.BalanceManager.UpdateBalance(ctx, &model.AssetBalance{
				AccountID: rec.AccountID,
				AssetID:   rec.AssetID,
				Quantity:  balance.Quantity + rec.Quantity,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// getAsset retrieves an asset by its symbol. If the asset is not found in the cache, it is retrieved from the database.
func (ingestor *TransactionIngestor) getAsset(assetSymbol string) (*model.Asset, error) {
	if assetFromCache, ok := ingestor.cache.assets.Get(assetSymbol); ok {
		log.Printf("found asset from cache: %v", assetSymbol)
		return assetFromCache.(*model.Asset), nil
	}

	log.Printf("retrieving asset from DB: %v", assetSymbol)
	asset, err := ingestor.AssetManager.FindAssetBySymbol(assetSymbol)
	if err != nil {
		return nil, err
	}

	ingestor.cache.assets.Set(assetSymbol, asset, cache.DefaultExpiration)

	return asset, nil
}
