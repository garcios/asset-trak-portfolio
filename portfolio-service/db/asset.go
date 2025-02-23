package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/model"
)

type AssetRepository struct {
	DB *sql.DB
}

func NewAssetRepository(db *sql.DB) *AssetRepository {
	return &AssetRepository{
		DB: db,
	}
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
