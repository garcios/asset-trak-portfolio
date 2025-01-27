package db

import (
	"database/sql"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/transaction-service/model"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(conn *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: conn}
}

func (r *TransactionRepository) AddTransaction(rec *model.Transaction) error {
	insertQuery := `INSERT INTO 
	   transaction (id,
		account_id,
		asset_id,
		transaction_type,
		transaction_date,
		quantity,
		price,
		currency_code) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.DB.Exec(insertQuery,
		rec.ID,
		rec.AccountID,
		rec.AssetID,
		rec.TransactionType,
		rec.TransactionDate,
		rec.Quantity,
		rec.Price,
		rec.CurrencyCode,
	)
	if err != nil {
		return fmt.Errorf("AddAsset: %v", err)
	}

	return nil
}

func (r *TransactionRepository) Truncate() error {
	_, err := r.DB.Exec("TRUNCATE transaction")
	if err != nil {
		return fmt.Errorf("truncate: %v", err)
	}

	return nil
}
