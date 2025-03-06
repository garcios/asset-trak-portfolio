package service

import (
	"encoding/csv"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/ingestion-service/db"
	"github.com/garcios/asset-trak-portfolio/ingestion-service/model"
	"github.com/patrickmn/go-cache"
	"log"
	"os"
	"strconv"
	"strings"
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

	log.Printf("%+v\n", ingestor.cfg)
	dirPath := ingestor.cfg.AssetPrice.DirPath

	// Read the directory contents.
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	// Iterate through the files and print their names.
	for _, file := range files {
		fmt.Println(file.Name())
		filePath := dirPath + "/" + file.Name()
		assetSymbol, err := getAssetSymbol(file.Name())
		if err != nil {
			return err
		}

		err = ingestor.processPrices(filePath, assetSymbol)
		if err != nil {
			return err
		}
	}

	return nil
}

// getAssetSymbol extracts the asset symbol from a given file name based on its format and separator.
// e.g. ASX.IVV.csv or US.AMZN.csv.
func getAssetSymbol(fileName string) (string, error) {
	// Remove the file extension.
	fileNameWithoutExt := strings.TrimSuffix(fileName, ".csv")

	// Split the string by the underscore.
	parts := strings.Split(fileNameWithoutExt, ".")

	// Check if the file name follows the expected format.
	if len(parts) != 2 {
		return "", fmt.Errorf("cannot parse asset symbol from file name: %s", fileName)
	}

	return parts[1], nil
}

func (ingestor *AssetPriceIngestor) processPrices(filePath string, assetSymbol string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Error opening file: %s\n", err.Error())
	}

	defer file.Close()

	// Create a new CSV reader.
	reader := csv.NewReader(file)

	// Read all records from the CSV file.
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("Error reading CSV data: %s\n", err.Error())
	}

	// Skip the header row (first row).
	if len(records) > 0 {
		records = records[1:]
	}

	for _, row := range records {
		// Ensure row has the expected number of fields.
		if len(row) < 6 {
			fmt.Printf("Skipping incomplete row: %v\n", row)
			continue
		}

		// Parse each field and handle errors.

		dateFormat := getDateFormat(row[0])

		// Parse the date string into a time.Time object.
		parsedDate, err := time.Parse(dateFormat, row[0])
		if err != nil {
			fmt.Printf("Error parsing date: %v\n", err)
			continue
		}

		closePrice, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			fmt.Printf("Error parsing Close value: %v\n", err)
			continue
		}

		asset, err := ingestor.getAsset(assetSymbol)
		if err != nil {
			return err
		}

		if asset == nil {
			return fmt.Errorf("asset %s not found", assetSymbol)
		}

		assetPrice := model.AssetPrice{
			AssetID:      asset.ID,
			Price:        closePrice,
			CurrencyCode: getCurrency(asset.MarketCode),
			TradeDate:    &parsedDate,
		}

		err = ingestor.assetPriceManager.AddAssetPrice(&assetPrice)
		if err != nil {
			return err
		}

	}

	return nil
}

func getCurrency(marketCode string) string {
	if marketCode == "ASX" {
		return "AUD"
	}

	return "USD"
}

func getDateFormat(dateStr string) string {
	if strings.Contains(dateStr, "-") {
		return "2006-01-02"
	}

	return "20060102"
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
