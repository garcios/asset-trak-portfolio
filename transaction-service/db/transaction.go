package db

import (
	"database/sql"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/transactions-service/model"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(conn *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: conn}
}

func (r *TransactionRepository) AddTransaction(rec *model.Transaction) error {
	_, err := r.DB.Exec("INSERT INTO transaction (id, symbol, name, market_code) VALUES (?, ?, ?,?)",
		rec.ID)
	if err != nil {
		return fmt.Errorf("AddAsset: %v", err)
	}

	return nil
}
