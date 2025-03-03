package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/asset-price-service/model"
	"log"
	"time"
)

type AssetPriceRepository struct {
	DB *sql.DB
}

func NewAssetPriceRepository(db *sql.DB) *AssetPriceRepository {
	return &AssetPriceRepository{
		DB: db,
	}
}

func (r *AssetPriceRepository) GetAssetPrice(
	assetID string,
	tradeDate time.Time,
) (*model.AssetPrice, error) {
	log.Println("db.GetAssetPrice")
	query := `SELECT price, currency_code
	           FROM asset_price
                WHERE asset_id = ?
                AND trade_date <= ?
                ORDER BY trade_date DESC LIMIT 1`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("GetExchangeRate: %v", err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(assetID, tradeDate)

	assetPrice := new(model.AssetPrice)
	if err := row.Scan(
		&assetPrice.Price,
		&assetPrice.CurrencyCode,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("GetAssetPrice: %v", err)
	}

	return assetPrice, nil
}
