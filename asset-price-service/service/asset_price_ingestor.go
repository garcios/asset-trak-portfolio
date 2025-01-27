package service

import (
	"fmt"
	"github.com/garcios/asset-trak-portfolio/asset-price-service/db"
	"github.com/garcios/asset-trak-portfolio/asset-price-service/model"
	"github.com/garcios/asset-trak-portfolio/lib/excel"
	"github.com/patrickmn/go-cache"
	"github.com/xuri/excelize/v2"
	"log"
	"time"
)

type allCache struct {
	assets *cache.Cache
}

const (
	defaultExpiration = 5 * time.Minute
	purgeTime         = 10 * time.Minute
)

type IAssetPriceManager interface {
	AddAssetPrice(rec *model.AssetPrice) error
	Truncate() error
}

type IAssetManager interface {
	FindAssetBySymbol(symbol string) (*model.Asset, error)
}

// verify interface compliance
var _ IAssetPriceManager = &db.AssetPriceRepository{}

type AssetPriceIngestor struct {
	assetPriceManager IAssetPriceManager
	assetManager      IAssetManager
	cfg               *Config
	symbolToCurrency  map[string]string
	cache             *allCache
}

func NewAssetPriceIngestor(
	assetPriceManager IAssetPriceManager,
	assetManager IAssetManager,
	cfg *Config,
) *AssetPriceIngestor {
	ac := cache.New(defaultExpiration, purgeTime)
	return &AssetPriceIngestor{
		assetPriceManager: assetPriceManager,
		assetManager:      assetManager,
		cfg:               cfg,
		symbolToCurrency:  make(map[string]string),
		cache:             &allCache{assets: ac},
	}
}

func (ingestor *AssetPriceIngestor) Truncate() error {
	err := ingestor.assetPriceManager.Truncate()
	if err != nil {
		return fmt.Errorf("truncate: %w", err)
	}

	return nil
}

func (ingestor *AssetPriceIngestor) ProcessAssetPrices() error {
	log.Println("Processing assets prices...")

	err := ingestor.loadCurrenTab()
	if err != nil {
		return err
	}

	log.Printf("%+v\n", ingestor.cfg)
	filePath := ingestor.cfg.FileInfo.Path
	tabs := ingestor.cfg.Asset.Symbols

	for _, tab := range tabs {
		err := ingestor.processPricesTab(filePath, tab)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ingestor *AssetPriceIngestor) processPricesTab(filePath string, tab string) error {
	rows, err := excel.GetRows(filePath, tab)
	if err != nil {
		return err
	}

	skipRows := ingestor.cfg.FileInfo.SkipRows

	var rowCount int
	for _, row := range rows {
		if rowCount < skipRows {
			rowCount++
			continue
		}

		tradeDate, err := getFloatAsDate(row[1])
		if err != nil {
			return err
		}

		price, err := getFloatValue(row[2])
		if err != nil {
			return err
		}

		currencyCode, ok := ingestor.symbolToCurrency[tab]
		if !ok {
			return fmt.Errorf("no such symbol to currency mapping")
		}

		asset, err := ingestor.getAsset(tab)
		if err != nil {
			return err
		}

		assetPrice := model.AssetPrice{
			AssetID:      asset.ID,
			Price:        price,
			CurrencyCode: currencyCode,
			TradeDate:    tradeDate,
		}

		err = ingestor.assetPriceManager.AddAssetPrice(&assetPrice)
		if err != nil {
			return err
		}

	}

	return nil
}

func getFloatAsDate(valueString string) (*time.Time, error) {
	floatValue, err := getFloatValue(valueString)
	if err != nil {
		return nil, err
	}

	dateValue, err := excelize.ExcelDateToTime(floatValue, false)
	if err != nil {
		return nil, err
	}

	return &dateValue, nil
}

func (ingestor *AssetPriceIngestor) loadCurrenTab() error {
	log.Println("Loading current tab...")
	rows, err := excel.GetRows(ingestor.cfg.FileInfo.Path, "current")
	if err != nil {
		return err
	}

	for _, row := range rows {
		if row[0] == "Symbol" {
			continue
		}
		ingestor.symbolToCurrency[row[1]] = row[3]
	}

	return nil
}

// getAsset retrieves an asset by its symbol. If the asset is not found in the cache, it is retrieved from the database.
func (ingestor *AssetPriceIngestor) getAsset(assetSymbol string) (*model.Asset, error) {
	if assetFromCache, ok := ingestor.cache.assets.Get(assetSymbol); ok {
		return assetFromCache.(*model.Asset), nil
	}

	asset, err := ingestor.assetManager.FindAssetBySymbol(assetSymbol)
	if err != nil {
		return nil, err
	}

	ingestor.cache.assets.Set(assetSymbol, asset, cache.DefaultExpiration)

	return asset, nil
}
