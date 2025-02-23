package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	pbc "github.com/garcios/asset-trak-portfolio/currency-service/proto"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/handler"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"log"
	"os"

	_ "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/etcd"

	stdlibTransactor "github.com/Thiht/transactor/stdlib"
	"github.com/garcios/asset-trak-portfolio/lib/mysql"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/db"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/service"

	pb "github.com/garcios/asset-trak-portfolio/portfolio-service/proto"
)

const (
	transactionIngestorProcessor = "transactionIngestor"
	truncateProcessor            = "truncate"
	portfolioServiceName         = "portfolio-service"
	currencyServiceName          = "currency-service"
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

	transactor, dbGetter := stdlibTransactor.NewTransactor(
		conn,
		stdlibTransactor.NestedTransactionsSavepoints,
	)

	ctx := context.Background()

	accountRepo := db.NewAccountRepository(conn)
	assetRepo := db.NewAssetRepository(conn)
	transactionRepo := db.NewTransactionRepository(dbGetter)
	portfolioRepo := db.NewAssetBalanceRepository(dbGetter)

	transactionIngestor := service.NewTransactionIngestor(
		transactionRepo,
		accountRepo,
		assetRepo,
		portfolioRepo,
		transactor,
		&cfg,
	)

	switch *processor {
	case transactionIngestorProcessor:
		err := transactionIngestor.ProcessTransactions(ctx)
		if err != nil {
			log.Fatalf("failed to process transaction ingestor: %v", err)
		}

		return
	case truncateProcessor:
		err := transactionIngestor.Truncate(ctx)
		if err != nil {
			log.Fatalf("failed to truncate transaction data: %v", err)
		}

		return
	default:
		log.Println("starting Transaction service...")
	}

	transactionSrv := micro.NewService(
		micro.Name(portfolioServiceName),
		micro.Version("latest"),
	)

	transactionSrv.Init()

	currencyService := pbc.NewCurrencyService(currencyServiceName, transactionSrv.Client())

	h := handler.New(currencyService, portfolioRepo, transactionRepo)

	err = pb.RegisterPortfolioHandler(transactionSrv.Server(), h)
	if err != nil {
		log.Fatalf("failed to register transaction handler: %v", err)
	}

	if err := transactionSrv.Run(); err != nil {
		logger.Fatal(err)
	}

	log.Println("Transaction service has started.")

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
