package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/currency-service/model"
	"time"
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

func (r *CurrencyRepository) GetExchangeRate(
	fromCurrency string,
	toCurrency string,
	tradeDate time.Time,
) (float64, error) {
	query := `SELECT exchange_rate
	           FROM currency_rate
                WHERE base_currency = ?
                AND target_currency = ?
                AND trade_date <= ?
                ORDER BY trade_date DESC LIMIT 1`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("GetExchangeRate: %v", err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(fromCurrency, toCurrency, tradeDate)

	var exchangeRate float64
	if err := row.Scan(&exchangeRate); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 1, nil
		}

		return 0, fmt.Errorf("GetExchangeRate: %v", err)
	}

	return 0, nil
}
