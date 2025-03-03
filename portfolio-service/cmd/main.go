package main

import (
	"context"
	"flag"
	"fmt"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	stdlibTransactor "github.com/Thiht/transactor/stdlib"
	_ "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/go-redis/redis/v8"

	pba "github.com/garcios/asset-trak-portfolio/asset-price-service/proto"
	pbc "github.com/garcios/asset-trak-portfolio/currency-service/proto"
	"github.com/garcios/asset-trak-portfolio/lib/mysql"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/db"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/handler"
	pb "github.com/garcios/asset-trak-portfolio/portfolio-service/proto"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/service"
)

const (
	tradesIngestorProcessor    = "tradesIngestor"
	dividendsIngestorProcessor = "dividendsIngestor"
	truncateProcessor          = "truncate"
	portfolioServiceName       = "portfolio-service"
	currencyServiceName        = "currency-service"
	assetPriceServiceName      = "asset-price-service"
)

func main() {
	processor := flag.String("processor", "", "the processor to run")
	accountID := flag.String("accountID", "", "the account ID to process")
	flag.Parse()

	cfg, err := readConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	// redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	// database
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

	dividendIngestor := service.NewDividendIngestor(
		transactionRepo,
		accountRepo,
		assetRepo,
		portfolioRepo,
		transactor,
		&cfg,
	)

	performnaceSvc := service.NewPerformanceService(rdb)

	switch *processor {
	case tradesIngestorProcessor:
		err := transactionIngestor.ProcessTrades(ctx, *accountID)
		if err != nil {
			log.Fatalf("failed to process transaction ingestor: %v", err)
		}

		return
	case dividendsIngestorProcessor:
		err := dividendIngestor.ProcessDividends(ctx, *accountID)
		if err != nil {
			log.Fatalf("failed to process dividend ingestor: %v", err)
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
	assetPriceService := pba.NewAssetPriceService(assetPriceServiceName, transactionSrv.Client())

	h := handler.New(currencyService, portfolioRepo, transactionRepo, performnaceSvc, assetPriceService)

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
