package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/garcios/asset-trak-portfolio/asset-price-service/db"
	"github.com/garcios/asset-trak-portfolio/asset-price-service/service"
	"github.com/garcios/asset-trak-portfolio/lib/mysql"
	"log"
)

const (
	assetPriceIngestorProcessor = "assetPriceIngestor"
	truncateProcessor           = "truncate"
)

func main() {
	processor := flag.String("processor", "", "the processor to run")
	flag.Parse()

	var cfg service.Config

	_, err := toml.DecodeFile("config.toml", &cfg)
	if err != nil {
		log.Fatalf("failed to load config.toml: %s", err)
	}

	conn, err := mysql.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	assetPriceRepo := db.NewAssetPriceRepository(conn)
	assetRepo := db.NewAssetRepository(conn)
	assetPriceIngestor := service.NewAssetPriceIngestor(assetPriceRepo, assetRepo, &cfg)

	switch *processor {
	case assetPriceIngestorProcessor:
		err := assetPriceIngestor.ProcessAssetPrices()
		if err != nil {
			log.Fatalf("failed to process asset price ingestor: %v", err)
		}

		return
	case truncateProcessor:
		err := assetPriceIngestor.Truncate()
		if err != nil {
			log.Fatalf("failed to truncate asset price ingestor: %v", err)
		}

		return
	default:
		log.Fatalf("unknown processor: %s", *processor)
	}

	//TODO: start gRPC service

	log.Println("Asset Price service is started.")
}
