package db

import (
	"context"
	"fmt"
	stdlibTransactor "github.com/Thiht/transactor/stdlib"
	"github.com/garcios/asset-trak-portfolio/transaction-service/model"
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
		rec.Price,
		rec.CurrencyCode,
	)
	if err != nil {
		return fmt.Errorf("AddAsset: %v", err)
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
