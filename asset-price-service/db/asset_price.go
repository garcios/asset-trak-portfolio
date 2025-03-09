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
		return nil, fmt.Errorf("GetAssetPrice: %v", err)
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

func (r *AssetPriceRepository) GetAssetPrices(assetID string, startDate, endDate string) ([]*model.AssetPrice, error) {
	query := `SELECT asset_id, price, currency_code, DATE(trade_date) FROM asset_price WHERE asset_id = ? AND trade_date BETWEEN ? AND ?`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("GetAssetPrice: %v", err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(assetID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("GetAssetPrice: %v", err)
	}
	defer rows.Close()

	var assetPrices []*model.AssetPrice
	var tradeDateStr string
	for rows.Next() {
		ap := &model.AssetPrice{}
		if err := rows.Scan(&ap.AssetID, &ap.Price, &ap.CurrencyCode, &tradeDateStr); err != nil {
			return nil, fmt.Errorf("failed to scan asset price: %w", err)
		}

		convertedDate, err := time.Parse("2006-01-02", tradeDateStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse trade_date: %v", err)
		}

		ap.TradeDate = &convertedDate

		assetPrices = append(assetPrices, ap)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return assetPrices, nil
}
