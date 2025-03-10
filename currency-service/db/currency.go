package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/currency-service/model"
	"log"
	"time"
)

var (
	NotFound = errors.New("not found")
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
			return 0, NotFound
		}

		return 0, fmt.Errorf("GetExchangeRate: %v", err)
	}

	return exchangeRate, nil
}

func (r *CurrencyRepository) GeExchangeRates(
	fromCurrency string,
	toCurrency string,
	startDate string,
	endDate string,
) ([]*model.CurrencyRate, error) {
	query := `
		SELECT base_currency, target_currency, exchange_rate, DATE(trade_date) 
		FROM currency_rate
		WHERE base_currency = ?
		  AND target_currency = ?
		  AND trade_date BETWEEN ? AND ?;
	`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("GetExchangeRate: %v", err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(fromCurrency, toCurrency, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to query exchange rates: %w", err)
	}
	defer rows.Close()

	exchangeRates := make([]*model.CurrencyRate, 0)
	var tradeDateStr string
	for rows.Next() {
		var rate model.CurrencyRate
		err := rows.Scan(&rate.BaseCurrency, &rate.TargetCurrency, &rate.ExchangeRate, &tradeDateStr)
		if err != nil {
			return nil, fmt.Errorf("failed to scan exchange rate: %w", err)
		}

		convertedDate, err := time.Parse("2006-01-02", tradeDateStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse trade_date: %v", err)
		}

		rate.TradeDate = &convertedDate

		exchangeRates = append(exchangeRates, &rate)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return exchangeRates, nil
}
