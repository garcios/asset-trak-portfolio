package service

import (
	"github.com/garcios/asset-trak-portfolio/ingestion-service/db"
	"github.com/garcios/asset-trak-portfolio/ingestion-service/model"
)

type ICurrencyManager interface {
	AddCurrencyRate(rec *model.CurrencyRate) error
	Truncate() error
}

type IAssetManager interface {
	AddAsset(rec *model.Asset) error
	AssetExists(symbol string, marketCode string) (bool, error)
	FindAssetBySymbol(symbol string) (*model.Asset, error)
	Truncate() error
}

var _ IAssetManager = &db.AssetRepository{}

// verify interface compliance
var _ ICurrencyManager = (*db.CurrencyRepository)(nil)
