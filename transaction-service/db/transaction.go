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
	_, err := r.DB.Exec("INSERT INTO "+
		"transaction (id,"+
		" account_id,"+
		" asset_id,"+
		" transaction_type,"+
		" transaction_date,"+
		" quantity,"+
		" price,"+
		" currency_code) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
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
