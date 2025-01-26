package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/transaction-service/model"
)

type AssetRepository struct {
	DB *sql.DB
}

func NewAssetRepository(db *sql.DB) *AssetRepository {
	return &AssetRepository{
		DB: db,
	}
}

func (r *AssetRepository) AddAsset(rec *model.Asset) error {
	_, err := r.DB.Exec("INSERT INTO asset (id, symbol, name, market_code) VALUES (?, ?, ?,?)",
		rec.ID, rec.Symbol, rec.Name, rec.MarketCode)
	if err != nil {
		return fmt.Errorf("AddAsset: %v", err)
	}

	return nil
}

func (r *AssetRepository) AssetExists(symbol string, marketCode string) (bool, error) {
	stmt, err := r.DB.Prepare("SELECT count(1) FROM asset WHERE symbol = ? AND market_code = ? LIMIT 1")
	if err != nil {
		return false, fmt.Errorf("AssetExists: %v", err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(symbol, marketCode)

	var count int
	if err := row.Scan(&count); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("AssetExists: %v", err)
	}

	if count == 1 {
		return true, nil
	}

	return false, nil
}

func (r *AssetRepository) FindAssetBySymbol(symbol string) (*model.Asset, error) {
	stmt, err := r.DB.Prepare("SELECT id, symbol, name, market_code FROM asset WHERE symbol = ? LIMIT 1")
	if err != nil {
		return nil, fmt.Errorf("FindAssetBySymbol: %v", err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(symbol)
	if row == nil {
		return nil, nil
	}

	asset := new(model.Asset)
	if err := row.Scan(&asset.ID, &asset.Symbol, &asset.Name, &asset.MarketCode); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("FindAssetBySymbol: %v", err)
	}

	return asset, nil
}
