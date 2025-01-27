package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/transaction-service/model"
)

type AssetBalanceRepository struct {
	DB *sql.DB
}

func NewAssetBalanceRepository(db *sql.DB) *AssetBalanceRepository {
	return &AssetBalanceRepository{
		DB: db,
	}
}

func (r *AssetBalanceRepository) AddBalance(rec *model.AssetBalance) error {
	insertQuery := "INSERT INTO asset_balance (account_id, asset_id, quantity) VALUES (?, ?, ?)"

	_, err := r.DB.Exec(insertQuery,
		rec.AccountID, rec.AssetID, rec.Quantity)
	if err != nil {
		return fmt.Errorf("AddBalance: %v", err)
	}

	return nil
}

func (r *AssetBalanceRepository) UpdateBalance(rec *model.AssetBalance) error {
	updateQuery := "UPDATE asset_balance SET quantity = ? WHERE account_id = ? AND asset_id = ?"

	_, err := r.DB.Exec(updateQuery,
		rec.Quantity, rec.AccountID, rec.AssetID)
	if err != nil {
		return fmt.Errorf("UpdateBalance: %v", err)
	}

	return nil
}

func (r *AssetBalanceRepository) GetBalance(accountID string, assetID string) (*model.AssetBalance, error) {
	stmt, err := r.DB.Prepare("SELECT account_id, asset_id, quantity FROM asset_balance WHERE account_id = ? AND asset_id = ? LIMIT 1")
	if err != nil {
		return nil, fmt.Errorf("BalanceExists: %v", err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(accountID, assetID)

	balance := new(model.AssetBalance)
	if err := row.Scan(&balance.AccountID, &balance.AssetID, &balance.Quantity); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("BalanceExists: %v", err)
	}

	return balance, nil
}

func (r *AssetBalanceRepository) Truncate() error {
	_, err := r.DB.Exec("TRUNCATE asset_balance")
	if err != nil {
		return fmt.Errorf("truncate: %v", err)
	}

	return nil
}
