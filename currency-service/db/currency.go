package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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

func (r *CurrencyRepository) GetExchangeRate(
	fromCurrency string,
	toCurrency string,
	tradeDate time.Time,
) (float64, error) {
	log.Println("db.GetExchangeRate")
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

	return exchangeRate, nil
}
