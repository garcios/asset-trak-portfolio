package db

import (
	"database/sql"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/ingestion-service/model"
)

type CurrencyRepository struct {
	DB *sql.DB
}

func NewCurrencyRepository(db *sql.DB) *CurrencyRepository {
	return &CurrencyRepository{
		DB: db,
	}
}

func (r *CurrencyRepository) AddCurrencyRate(rec *model.CurrencyRate) error {
	query := "INSERT INTO currency_rate (base_currency, target_currency, exchange_rate, trade_date) VALUES (?, ?, ?,?)"

	_, err := r.DB.Exec(query,
		rec.BaseCurrency, rec.TargetCurrency, rec.ExchangeRate, rec.TradeDate)
	if err != nil {
		return fmt.Errorf("AddCurrencyRate: %v", err)
	}

	return nil
}

func (r *CurrencyRepository) Truncate() error {
	_, err := r.DB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	if err != nil {
		return fmt.Errorf("truncate: %v", err)
	}

	_, err = r.DB.Exec("TRUNCATE currency_rate")
	if err != nil {
		return fmt.Errorf("truncate: %v", err)
	}

	_, err = r.DB.Exec("SET FOREIGN_KEY_CHECKS = 1")
	if err != nil {
		return fmt.Errorf("truncate: %v", err)
	}

	return nil
}
