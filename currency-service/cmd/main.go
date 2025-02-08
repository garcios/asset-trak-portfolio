package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/garcios/asset-trak-portfolio/currency-service/db"
	"github.com/garcios/asset-trak-portfolio/currency-service/handler"
	pb "github.com/garcios/asset-trak-portfolio/currency-service/proto"
	"github.com/garcios/asset-trak-portfolio/currency-service/service"
	"github.com/garcios/asset-trak-portfolio/lib/mysql"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"log"
	"os"
)

const (
	currencyRateIngestorProcessor = "currencyRateIngestor"
	truncateProcessor             = "truncate"
	serviceName                   = "currency-service"
)

func main() {
	processor := flag.String("processor", "", "the processor to run")
	flag.Parse()

	var cfg service.Config

	configDir := os.Getenv("CONFIG_DIR")
	if configDir == "" {
		configDir = "./"
	}

	configPath := configDir + "config.toml"

	_, err := toml.DecodeFile(configPath, &cfg)
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

	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
	)

	srv.Init()

	h := handler.New(currencyRepo)

	err = pb.RegisterCurrencyHandler(srv.Server(), h)
	if err != nil {
		log.Fatalf("failed to register currency handler: %v", err)
	}

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}

	log.Println("Currency service is started.")
}
