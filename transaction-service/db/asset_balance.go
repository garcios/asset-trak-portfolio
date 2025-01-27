package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	stdlibTransactor "github.com/Thiht/transactor/stdlib"
	"github.com/garcios/asset-trak-portfolio/transaction-service/model"
)

type AssetBalanceRepository struct {
	dbGetter stdlibTransactor.DBGetter
}

func NewAssetBalanceRepository(dbGetter stdlibTransactor.DBGetter) *AssetBalanceRepository {
	return &AssetBalanceRepository{
		dbGetter: dbGetter,
	}
}

func (r *AssetBalanceRepository) AddBalance(ctx context.Context, rec *model.AssetBalance) error {
	insertQuery := "INSERT INTO asset_balance (account_id, asset_id, quantity) VALUES (?, ?, ?)"

	_, err := r.dbGetter(ctx).Exec(insertQuery,
		rec.AccountID, rec.AssetID, rec.Quantity)
	if err != nil {
		return fmt.Errorf("AddBalance: %v", err)
	}

	return nil
}

func (r *AssetBalanceRepository) UpdateBalance(ctx context.Context, rec *model.AssetBalance) error {
	updateQuery := "UPDATE asset_balance SET quantity = ? WHERE account_id = ? AND asset_id = ?"

	_, err := r.dbGetter(ctx).Exec(updateQuery,
		rec.Quantity, rec.AccountID, rec.AssetID)
	if err != nil {
		return fmt.Errorf("UpdateBalance: %v", err)
	}

	return nil
}

func (r *AssetBalanceRepository) GetBalance(ctx context.Context, accountID string, assetID string) (*model.AssetBalance, error) {
	query := "SELECT account_id, asset_id, quantity FROM asset_balance WHERE account_id = ? AND asset_id = ?"
	if stdlibTransactor.IsWithinTransaction(ctx) {
		query += ` FOR UPDATE`
	}

	stmt, err := r.dbGetter(ctx).Prepare(query)
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

func (r *AssetBalanceRepository) Truncate(ctx context.Context) error {
	_, err := r.dbGetter(ctx).Exec("TRUNCATE asset_balance")
	if err != nil {
		return fmt.Errorf("truncate: %v", err)
	}

	return nil
}
