package db

import (
	"database/sql"
)

type AssetPriceRepository struct {
	DB *sql.DB
}

func NewAssetPriceRepository(db *sql.DB) *AssetPriceRepository {
	return &AssetPriceRepository{
		DB: db,
	}
}
