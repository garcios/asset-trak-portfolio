package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/garcios/asset-trak-portfolio/currency-service/db"
	"github.com/garcios/asset-trak-portfolio/currency-service/service"
	"github.com/garcios/asset-trak-portfolio/lib/mysql"
	"log"
)

const (
	currencyRateIngestorProcessor = "currencyRateIngestor"
	truncateProcessor             = "truncate"
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

	currencyRepo := db.NewCurrencyRepository(conn)
	currencyIngestor := service.NewCurrencyIngestor(currencyRepo, &cfg)

	switch *processor {
	case currencyRateIngestorProcessor:
		err := currencyIngestor.ProcessCurrencyRates()
		if err != nil {
			log.Fatalf("failed to process currency ingestor: %v", err)
		}

		return
	case truncateProcessor:
		err := currencyIngestor.Truncate()
		if err != nil {
			log.Fatalf("failed to truncate currency ingestor: %v", err)
		}

		return
	default:
		log.Println("starting currency service...")
	}

	//TODO: start gRPC service

	log.Println("Currency service is started.")
}
