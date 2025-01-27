package service

import (
	"fmt"
	"github.com/garcios/asset-trak-portfolio/lib/excel"
	"github.com/garcios/asset-trak-portfolio/transaction-service/db"
	"github.com/garcios/asset-trak-portfolio/transaction-service/model"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"log"
	"strconv"
	"strings"
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
	AddTransaction(rec *model.Transaction) error
}

type IAccountManager interface {
	FindAccountByID(id string) (*model.Account, error)
}

// verify interface compliance
var _ ITransactionManager = &db.TransactionRepository{}
var _ IAccountManager = &db.AccountRepository{}

type TransactionIngestor struct {
	TransactionManager ITransactionManager
	AccountManager     IAccountManager
	AssetManager       IAssetManager
	cache              *allCache
}

func NewTransactionIngestor(
	tm ITransactionManager,
	am IAccountManager,
	astm IAssetManager,
) *TransactionIngestor {
	ac := cache.New(defaultExpiration, purgeTime)
	return &TransactionIngestor{
		TransactionManager: tm,
		AccountManager:     am,
		AssetManager:       astm,
		cache:              &allCache{assets: ac},
	}
}

func (ingestor *TransactionIngestor) ProcessTransactions(
	filePath,
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

		err = ingestor.addTransaction(transaction)
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

	transactionDate, err := getDateValue(row[3])
	if err != nil {
		displayRow(row)
		return nil, fmt.Errorf("unable to process transaction date with value: %s", row[3])
	}

	quantity, err := getFloatValue(row[5])
	if err != nil {
		displayRow(row)
		return nil, fmt.Errorf("unable to process quantity with value: %s", row[5])
	}

	price, err := getFloatValue(row[6])
	if err != nil {
		displayRow(row)
		return nil, fmt.Errorf("unable to process price with value: %s", row[6])
	}

	return &model.Transaction{
		AssetID:         asset.ID,
		TransactionType: getStringValue(row[4]),
		TransactionDate: transactionDate,
		Quantity:        int(quantity),
		Price:           price,
		CurrencyCode:    row[7],
	}, nil
}

func (ingestor *TransactionIngestor) addTransaction(rec *model.Transaction) error {

	// TODO:
	// Insert asset if it does not exist yet.

	// The following steps should be atomic.
	// 1. insert into transaction table
	// 2. insert/update asset balance
	//     - if asset for account_id already exists, insert; otherwise update balance.

	err := ingestor.TransactionManager.AddTransaction(rec)
	if err != nil {
		return err
	}

	return nil
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

func getDateValue(dateString string) (*time.Time, error) {
	if dateString == "" {
		return nil, nil
	}

	dateValue, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return nil, err
	}

	return &dateValue, nil
}

func getFloatValue(valueString string) (float64, error) {
	if valueString == "" {
		return 0, nil
	}

	value, err := strconv.ParseFloat(normalizeNumber(valueString), 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func normalizeNumber(s string) string {
	return strings.Replace(s, ",", "", -1)
}

func getStringValue(valueString string) string {
	if valueString == "" {
		return ""
	}

	return strings.ToUpper(strings.TrimSpace(valueString))
}

func displayRow(row []string) {
	log.Printf("row: %v", row)
}
