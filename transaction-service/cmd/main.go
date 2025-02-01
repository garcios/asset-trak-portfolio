package main

import (
	"context"
	"flag"
	"github.com/garcios/asset-trak-portfolio/transaction-service/handler"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"log"

	_ "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/etcd"

	stdlibTransactor "github.com/Thiht/transactor/stdlib"
	"github.com/garcios/asset-trak-portfolio/lib/mysql"
	"github.com/garcios/asset-trak-portfolio/transaction-service/db"
	"github.com/garcios/asset-trak-portfolio/transaction-service/service"

	pb "github.com/garcios/asset-trak-portfolio/transaction-service/proto"
)

const (
	assetIngestorProcessor       = "assetIngestor"
	transactionIngestorProcessor = "transactionIngestor"
	truncateProcessor            = "truncate"
	skipRows                     = 3
	tabName                      = "Combined"
	filePath                     = "data/AllTradesReport.xlsx"
	accountID                    = "eb08df3c-958d-4ae8-b3ae-41ec04418786"
	serviceName                  = "transaction-service"
)

func main() {
	processor := flag.String("processor", "", "the processor to run")
	flag.Parse()

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
	balanceRepo := db.NewAssetBalanceRepository(dbGetter)

	assetIngestor := service.NewAssetIngestor(assetRepo)
	transactionIngestor := service.NewTransactionIngestor(
		transactionRepo,
		accountRepo,
		assetRepo,
		balanceRepo,
		transactor,
	)

	switch *processor {
	case assetIngestorProcessor:
		err := assetIngestor.ProcessAssets(filePath, tabName, skipRows)
		if err != nil {
			log.Fatalf("failed to process asset ingestor: %v", err)
		}

		return
	case transactionIngestorProcessor:
		err := transactionIngestor.ProcessTransactions(ctx, filePath, tabName, skipRows, accountID)
		if err != nil {
			log.Fatalf("failed to process transaction ingestor: %v", err)
		}

		return
	case truncateProcessor:
		err := transactionIngestor.Truncate(ctx)
		if err != nil {
			log.Fatalf("failed to truncate transaction data: %v", err)
		}

		err = assetIngestor.Truncate()
		if err != nil {
			log.Fatalf("failed to truncate asset data: %v", err)
		}

		return
	default:
		log.Println("starting Transaction service...")
	}

	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
	)

	srv.Init()

	h := handler.New(srv.Client())

	pb.RegisterTransactionHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}

	log.Println("Transaction service has started.")

}
