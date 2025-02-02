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
		return nil, fmt.Errorf("GetBalance: %v", err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(accountID, assetID)

	balance := new(model.AssetBalance)
	if err := row.Scan(&balance.AccountID, &balance.AssetID, &balance.Quantity); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("GetBalance: %v", err)
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

func (r *AssetBalanceRepository) GetBalanceSummary(
	ctx context.Context,
	accountID string,
) ([]*model.BalanceSummary, error) {
	query := `SELECT a.symbol, a.name, b.quantity, p.price, p.currency_code, a.market_code
              FROM asset_balance b JOIN asset a ON b.asset_id = a.id
              LEFT JOIN (
                 SELECT asset_id, price, trade_date, currency_code,
                        ROW_NUMBER() OVER (PARTITION BY asset_id ORDER BY trade_date DESC) as rn
                 FROM asset_price WHERE trade_date <= NOW()) p
                  ON p.asset_id = a.id AND p.rn = 1 
              WHERE b.quantity > 0
                AND b.account_id = ?`

	stmt, err := r.dbGetter(ctx).Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("GetBalanceSummary: %v", err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(accountID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("GetBalanceSummary: %v", err)
	}

	summary := make([]*model.BalanceSummary, 0)
	for rows.Next() {
		var summaryItem model.BalanceSummary
		if err := rows.Scan(
			&summaryItem.AssetSymbol,
			&summaryItem.AssetName,
			&summaryItem.Quantity,
			&summaryItem.Price,
			&summaryItem.CurrencyCode,
			&summaryItem.MarketCode,
		); err != nil {
			return nil, fmt.Errorf("GetBalanceSummary: %v", err)
		}
		summary = append(summary, &summaryItem)
	}

	return summary, nil

}
