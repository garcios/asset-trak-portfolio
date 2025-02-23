package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/garcios/asset-trak-portfolio/ingestion-service/db"
	"github.com/garcios/asset-trak-portfolio/ingestion-service/service"
	"github.com/garcios/asset-trak-portfolio/lib/mysql"
	"log"
	"os"
)

const (
	ingestAssetProcessor   = "ingestAsset"
	truncateAssetProcessor = "truncateAsset"

	ingestAssetPriceProcessor   = "ingestAssetPrice"
	truncateAssetPriceProcessor = "truncateAssetPrice"

	ingestCurrencyRatesProcessor   = "ingestCurrencyRates"
	truncateCurrencyRatesProcessor = "truncateCurrencyRates"
)

func main() {
	processor := flag.String("processor", "", "the processor to run")
	flag.Parse()

	cfg, err := readConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	conn, err := mysql.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// asset
	assetRepo := db.NewAssetRepository(conn)
	assetIngestor := service.NewAssetIngestor(assetRepo, &cfg)

	// asset prices
	assetPriceRepo := db.NewAssetPriceRepository(conn)
	assetPriceIngestor := service.NewAssetPriceIngestor(assetPriceRepo, assetRepo, &cfg)

	// currency
	currencyRepo := db.NewCurrencyRepository(conn)
	currencyIngestor := service.NewCurrencyIngestor(currencyRepo, &cfg)

	switch *processor {
	case ingestAssetProcessor:
		err := assetIngestor.ProcessAssets()
		if err != nil {
			log.Fatalf("failed to process asset ingestor: %v", err)
		}

		return
	case truncateAssetProcessor:
		err := assetIngestor.Truncate()
		if err != nil {
			log.Fatalf("failed to truncate assets: %v", err)
		}

		return
	case ingestAssetPriceProcessor:
		err := assetPriceIngestor.ProcessAssetPrices()
		if err != nil {
			log.Fatalf("failed to process asset price ingestor: %v", err)
		}

		return
	case truncateAssetPriceProcessor:
		err := assetPriceIngestor.Truncate()
		if err != nil {
			log.Fatalf("failed to truncate asset prices: %v", err)
		}

		return
	case ingestCurrencyRatesProcessor:
		err := currencyIngestor.ProcessCurrencyRates()
		if err != nil {
			log.Fatalf("failed to process currency rates ingestor: %v", err)
		}

		return
	case truncateCurrencyRatesProcessor:
		err := currencyIngestor.Truncate()
		if err != nil {
			log.Fatalf("failed to truncate currency rates: %v", err)
		}

		return
	default:
		log.Println("Ingestor service is being started...")
	}

	// TODO: setup starting this service as background process for processing external API market data.

}

func readConfig() (service.Config, error) {
	var cfg service.Config
	configDir := os.Getenv("CONFIG_DIR")
	if configDir == "" {
		configDir = "./"
	}

	configPath := configDir + "config.toml"

	_, err := toml.DecodeFile(configPath, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("failed to load config.toml: %s", err)
	}

	return cfg, nil
}
