package service

import (
	"context"
	"fmt"
	"github.com/Thiht/transactor"
	"github.com/garcios/asset-trak-portfolio/lib/excel"
	lib "github.com/garcios/asset-trak-portfolio/lib/typesutils"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/model"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"log"
	"strings"
)

const (
	TransactionTypeDividend = "DIVIDEND"
)

type DividendIngestor struct {
	TransactionManager ITransactionManager
	AccountManager     IAccountManager
	AssetManager       IAssetManager
	BalanceManager     IBalanceManager
	Transactor         transactor.Transactor
	cache              *allCache
	cfg                *Config
}

func NewDividendIngestor(
	tm ITransactionManager,
	am IAccountManager,
	astm IAssetManager,
	bm IBalanceManager,
	tr transactor.Transactor,
	cfg *Config,
) *TransactionIngestor {
	ac := cache.New(defaultExpiration, purgeTime)
	return &TransactionIngestor{
		TransactionManager: tm,
		AccountManager:     am,
		AssetManager:       astm,
		BalanceManager:     bm,
		Transactor:         tr,
		cfg:                cfg,
		cache:              &allCache{assets: ac},
	}
}

func (ingestor *TransactionIngestor) ProcessDividends(
	ctx context.Context,
	accountID string,
) error {
	log.Println("Processing dividends...")

	dividendCfg := ingestor.cfg.Dividends

	// verify account
	account, err := ingestor.AccountManager.FindAccountByID(accountID)
	if err != nil {
		return err
	}

	if account == nil {
		return fmt.Errorf("account with ID %s not found", accountID)
	}

	err = ingestor.domesticDividend(ctx, dividendCfg, accountID)
	if err != nil {
		return err
	}

	err = ingestor.foreignDividend(ctx, dividendCfg, accountID)
	if err != nil {
		return err
	}

	return nil

}

func (ingestor *TransactionIngestor) domesticDividend(
	ctx context.Context,
	dividendCfg Dividends,
	accountID string,
) error {
	rows, err := excel.GetRows(
		dividendCfg.Path,
		dividendCfg.TabNameDomestic)
	if err != nil {
		return err
	}

	var rowCount int
	for _, row := range rows {
		if rowCount < dividendCfg.SkipRowsDomestic {
			rowCount++
			continue
		}

		if len(row) < 2 || row[0] == "Total" || row[0] == "Trust Income" || row[0] == "Code" {
			continue
		}

		if strings.Contains(row[0], "Grand Total") {
			break
		}

		transaction, err := ingestor.mapDomesticDividendColumnsToTransaction(row)
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

func (ingestor *TransactionIngestor) foreignDividend(
	ctx context.Context,
	dividendCfg Dividends,
	accountID string,
) error {
	rows, err := excel.GetRows(
		dividendCfg.Path,
		dividendCfg.TabNameForeign)
	if err != nil {
		return err
	}

	var rowCount int
	for _, row := range rows {
		if rowCount < dividendCfg.SkipRowsForeign {
			rowCount++
			continue
		}

		if len(row) < 2 || row[0] == "Total" {
			break
		}

		transaction, err := ingestor.mapForeignDividendColumnsToTransaction(row)
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

func (ingestor *TransactionIngestor) addDividend(ctx context.Context, rec *model.Transaction) error {
	return ingestor.Transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		err := ingestor.TransactionManager.AddTransaction(ctx, rec)
		if err != nil {
			return err
		}

		return nil
	})
}

func (ingestor *TransactionIngestor) mapDomesticDividendColumnsToTransaction(row []string) (*model.Transaction, error) {
	assetSymbol, err := getAssetSymbol(row[0])
	if err != nil {
		return nil, err
	}

	asset, err := ingestor.getAsset(assetSymbol)
	if err != nil {
		displayRow(row)
		return nil, fmt.Errorf("unable to retrieve asset with symbol: %s", row[0])
	}

	transactionDate, err := lib.GetDateValue(row[2], "02/01/2006")
	if err != nil {
		displayRow(row)
		return nil, fmt.Errorf("unable to process transaction date with value: %s", row[2])
	}

	amountCash, err := lib.GetFloatValue(row[3])
	if err != nil {
		displayRow(row)
		return nil, fmt.Errorf("unable to process amount cash with value: %s", row[3])
	}

	return &model.Transaction{
		AssetID:                 asset.ID,
		TransactionType:         TransactionTypeDividend,
		TransactionDate:         transactionDate,
		AmountCash:              amountCash,
		AmountCurrencyCode:      "AUD",
		ExchangeRate:            1.0,
		WithheldTaxAmount:       0.0,
		WithheldTaxCurrencyCode: "AUD",
	}, nil
}

func (ingestor *TransactionIngestor) mapForeignDividendColumnsToTransaction(row []string) (*model.Transaction, error) {
	assetSymbol, err := getAssetSymbol(row[0])
	if err != nil {
		return nil, err
	}

	asset, err := ingestor.getAsset(assetSymbol)
	if err != nil {
		displayRow(row)
		return nil, fmt.Errorf("unable to retrieve asset with symbol: %s", row[0])
	}

	transactionDate, err := lib.GetDateValue(row[2], "02/01/2006")
	if err != nil {
		displayRow(row)
		return nil, fmt.Errorf("unable to process transaction date with value: %s", row[2])
	}

	exchangeRate, err := lib.GetFloatValue(row[3])
	if err != nil {
		displayRow(row)
		return nil, fmt.Errorf("unable to process amount cash with value: %s", row[3])
	}

	amountCash, err := lib.GetFloatValue(row[7])
	if err != nil {
		displayRow(row)
		return nil, fmt.Errorf("unable to process amount cash with value: %s", row[7])
	}

	withheldAmount, err := lib.GetFloatValue(row[6])
	if err != nil {
		displayRow(row)
		return nil, fmt.Errorf("unable to process amount cash with value: %s", row[6])
	}

	return &model.Transaction{
		AssetID:                 asset.ID,
		TransactionType:         TransactionTypeDividend,
		TransactionDate:         transactionDate,
		AmountCash:              amountCash,
		AmountCurrencyCode:      strings.ToUpper(row[4]),
		ExchangeRate:            exchangeRate,
		WithheldTaxAmount:       withheldAmount,
		WithheldTaxCurrencyCode: strings.ToUpper(row[4]),
	}, nil
}

func getAssetSymbol(column string) (string, error) {
	assetSymbolAndMarket := strings.Split(column, ".")
	if len(assetSymbolAndMarket) != 2 {
		return "", fmt.Errorf("row contains less than expected fields: %d", len(assetSymbolAndMarket))
	}

	return assetSymbolAndMarket[0], nil
}
