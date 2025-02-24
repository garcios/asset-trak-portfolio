package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	stdlibTransactor "github.com/Thiht/transactor/stdlib"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/model"
	"strings"
	"time"
)

type TransactionRepository struct {
	dbGetter stdlibTransactor.DBGetter
}

type TransactionFilter struct {
	AccountID        string
	AssetID          string
	StartDate        string
	EndDate          string
	transactionTypes []string
}

func NewTransactionRepository(dbGetter stdlibTransactor.DBGetter) *TransactionRepository {
	return &TransactionRepository{dbGetter: dbGetter}
}

func (r *TransactionRepository) AddTransaction(ctx context.Context, rec *model.Transaction) error {
	insertQuery := `INSERT INTO 
	   transaction (id,
		account_id,
		asset_id,
		transaction_type,
		transaction_date,
		quantity,
		trade_price,
		trade_price_currency_code,
	    brokerage_fee,
	    fee_currency_code,
	    amount_cash,
	    amount_currency_code,            
	    exchange_rate,
	    withheld_tax_amount,
	    withheld_tax_currency_code) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.dbGetter(ctx).Exec(insertQuery,
		rec.ID,
		rec.AccountID,
		rec.AssetID,
		rec.TransactionType,
		rec.TransactionDate,
		rec.Quantity,
		rec.TradePrice,
		rec.TradePriceCurrencyCode,
		rec.BrokerageFee,
		rec.FeeCurrencyCode,
		rec.AmountCash,
		rec.AmountCurrencyCode,
		rec.ExchangeRate,
		rec.WithheldTaxAmount,
		rec.WithheldTaxCurrencyCode,
	)
	if err != nil {
		return fmt.Errorf("AddTransaction: %v", err)
	}

	return nil
}

func (r *TransactionRepository) Truncate(ctx context.Context) error {
	_, err := r.dbGetter(ctx).Exec("TRUNCATE transaction")
	if err != nil {
		return fmt.Errorf("truncate: %v", err)
	}

	return nil
}

func (r *TransactionRepository) GetTransactions(
	ctx context.Context,
	filter TransactionFilter,
) ([]*model.Transaction, error) {
	query := `SELECT id,
       account_id,
       asset_id,
       transaction_type,
       transaction_date,
       quantity,
       trade_price, 
       trade_price_currency_code, 
       brokerage_fee,
       fee_currency_code,
       amount_cash,
       amount_currency_code,
       exchange_rate,
       withheld_tax_amount,
       withheld_tax_currency_code
       FROM transaction WHERE account_id = ?`
	args := []interface{}{filter.AccountID}

	if filter.AssetID != "" {
		query += ` AND asset_id = ?`
		args = append(args, filter.AssetID)
	}

	if filter.StartDate != "" {
		query += ` AND transaction_date >= ?`
		args = append(args, filter.StartDate)
	}

	if filter.EndDate != "" {
		query += ` AND transaction_date <= ?`
		args = append(args, filter.EndDate)
	}

	if len(filter.transactionTypes) > 0 {
		// Generate placeholders for the number of transaction types
		placeholders := make([]string, len(filter.transactionTypes))
		for i := range filter.transactionTypes {
			placeholders[i] = "?"
		}

		// Join the placeholders with commas
		query += ` AND transaction_type IN (` + strings.Join(placeholders, ",") + `)`

		// Append all transaction type values to the args slice
		for _, transactionType := range filter.transactionTypes {
			args = append(args, transactionType)
		}
	}

	query += ` ORDER BY transaction_date ASC`

	if stdlibTransactor.IsWithinTransaction(ctx) {
		query += ` FOR UPDATE`
	}

	stmt, err := r.dbGetter(ctx).Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("GetTransactions: %v", err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
	}
	transactions := make([]*model.Transaction, 0)

	var transactionDateStr string
	for rows.Next() {
		var transaction model.Transaction
		if err := rows.Scan(
			&transaction.ID,
			&transaction.AccountID,
			&transaction.AssetID,
			&transaction.TransactionType,
			&transactionDateStr,
			&transaction.Quantity,
			&transaction.TradePrice,
			&transaction.TradePriceCurrencyCode,
			&transaction.BrokerageFee,
			&transaction.FeeCurrencyCode,
			&transaction.AmountCash,
			&transaction.AmountCurrencyCode,
			&transaction.ExchangeRate,
			&transaction.WithheldTaxAmount,
			&transaction.WithheldTaxCurrencyCode,
		); err != nil {
			return nil, fmt.Errorf("GetTransactions: %v", err)
		}

		convertedDate, err := time.Parse("2006-01-02", transactionDateStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse transaction_date: %v", err)
		}

		transaction.TransactionDate = &convertedDate

		transactions = append(transactions, &transaction)
	}

	return transactions, nil
}
