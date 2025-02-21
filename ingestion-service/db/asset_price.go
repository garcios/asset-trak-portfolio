package db

import (
	"database/sql"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/asset-price-service/model"
)

type AssetPriceRepository struct {
	DB *sql.DB
}

func NewAssetPriceRepository(db *sql.DB) *AssetPriceRepository {
	return &AssetPriceRepository{
		DB: db,
	}
}

func (r *AssetPriceRepository) AddAssetPrice(rec *model.AssetPrice) error {
	query := "INSERT INTO asset_price (asset_id, price, currency_code, trade_date) VALUES (?, ?, ?,?)"

	_, err := r.DB.Exec(query,
		rec.AssetID, rec.Price, rec.CurrencyCode, rec.TradeDate)
	if err != nil {
		return fmt.Errorf("AddAssetPrice: %v", err)
	}

	return nil
}

func (r *AssetPriceRepository) Truncate() error {
	_, err := r.DB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	if err != nil {
		return fmt.Errorf("truncate: %v", err)
	}

	_, err = r.DB.Exec("TRUNCATE asset_price")
	if err != nil {
		return fmt.Errorf("truncate: %v", err)
	}

	_, err = r.DB.Exec("SET FOREIGN_KEY_CHECKS = 1")
	if err != nil {
		return fmt.Errorf("truncate: %v", err)
	}

	return nil
}
