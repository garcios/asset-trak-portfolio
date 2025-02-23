package service

import (
	"fmt"
	"github.com/garcios/asset-trak-portfolio/ingestion-service/model"
	"github.com/garcios/asset-trak-portfolio/lib/excel"
	"github.com/garcios/asset-trak-portfolio/lib/typesutils"
	"log"
)

const (
	baseCurrency   = "USD"
	targetCurrency = "AUD"
)

type CurrencyIngestor struct {
	currencyManager ICurrencyManager
	cfg             *Config
}

func NewCurrencyIngestor(
	currencyManager ICurrencyManager,
	cfg *Config,
) *CurrencyIngestor {
	return &CurrencyIngestor{
		currencyManager: currencyManager,
		cfg:             cfg,
	}
}

func (ingestor *CurrencyIngestor) Truncate() error {
	err := ingestor.currencyManager.Truncate()
	if err != nil {
		return fmt.Errorf("truncate: %w", err)
	}

	return nil
}

func (ingestor *CurrencyIngestor) ProcessCurrencyRates() error {
	log.Println("Processing currency rates...")

	log.Printf("%+v\n", ingestor.cfg)
	filePath := ingestor.cfg.CurrencyRate.Path

	err := ingestor.processCurrencyRates(filePath, "price-history")
	if err != nil {
		return fmt.Errorf("processPricesTab: %w", err)
	}

	return nil
}

func (ingestor *CurrencyIngestor) processCurrencyRates(filePath string, tab string) error {
	rows, err := excel.GetRows(filePath, tab)
	if err != nil {
		return err
	}

	skipRows := ingestor.cfg.CurrencyRate.SkipRows

	var rowCount int
	for _, row := range rows {
		if rowCount < skipRows {
			rowCount++
			continue
		}

		tradeDate, err := typesutils.GetFloatAsDate(row[1])
		if err != nil {
			return err
		}

		rate, err := typesutils.GetFloatValue(row[2])
		if err != nil {
			return err
		}

		currencyRate := model.CurrencyRate{
			BaseCurrency:   baseCurrency,
			TargetCurrency: targetCurrency,
			ExchangeRate:   rate,
			TradeDate:      tradeDate,
		}

		err = ingestor.currencyManager.AddCurrencyRate(&currencyRate)
		if err != nil {
			return err
		}

	}

	return nil
}
