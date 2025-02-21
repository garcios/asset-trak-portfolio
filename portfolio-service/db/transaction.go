package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	stdlibTransactor "github.com/Thiht/transactor/stdlib"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/model"
	"time"
)

type TransactionRepository struct {
	dbGetter stdlibTransactor.DBGetter
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
		price,
		currency_code) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.dbGetter(ctx).Exec(insertQuery,
		rec.ID,
		rec.AccountID,
		rec.AssetID,
		rec.TransactionType,
		rec.TransactionDate,
		rec.Quantity,
		rec.TradePrice,
		rec.CurrencyCode,
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
	accountID string,
	assetID string,
	startDate string,
	endDate string,
) ([]*model.Transaction, error) {
	query := `SELECT id, account_id, asset_id, transaction_type, transaction_date, quantity, price, currency_code FROM transaction WHERE account_id = ?`
	args := []interface{}{accountID}

	if assetID != "" {
		query += ` AND asset_id = ?`
		args = append(args, assetID)
	}

	if startDate != "" {
		query += ` AND transaction_date >= ?`
		args = append(args, startDate)
	}

	if endDate != "" {
		query += ` AND transaction_date <= ?`
		args = append(args, endDate)
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

	var dateStr string
	for rows.Next() {
		var transaction model.Transaction
		if err := rows.Scan(
			&transaction.ID,
			&transaction.AccountID,
			&transaction.AssetID,
			&transaction.TransactionType,
			&dateStr,
			&transaction.Quantity,
			&transaction.TradePrice,
			&transaction.CurrencyCode,
		); err != nil {
			return nil, fmt.Errorf("GetTransactions: %v", err)
		}

		convertedDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse transaction_date: %v", err)
		}

		transaction.TransactionDate = &convertedDate

		transactions = append(transactions, &transaction)
	}

	return transactions, nil
}
