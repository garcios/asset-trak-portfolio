package main

import (
	"flag"
	"github.com/garcios/asset-trak-portfolio/lib/mysql"
	"github.com/garcios/asset-trak-portfolio/transactions-service/db"
	"github.com/garcios/asset-trak-portfolio/transactions-service/service"
	"log"
)

const (
	assetIngestor       = "assetIngestor"
	transactionIngestor = "transactionIngestor"
	skipRows            = 3
	tabName             = "Combined"
	filePath            = "data/AllTradesReport.xlsx"
	accountID           = "eb08df3c-958d-4ae8-b3ae-41ec04418786"
)

func main() {
	processor := flag.String("processor", "", "a string")
	flag.Parse()

	conn, err := mysql.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	assetRepo := db.NewAssetRepository(conn)
	transactionRepo := db.NewTransactionRepository(conn)
	accountRepo := db.NewAccountRepository(conn)

	if *processor == assetIngestor {
		assetIngestor := service.NewAssetIngestor(assetRepo)

		err = assetIngestor.ProcessAssets(filePath, tabName, skipRows)
		if err != nil {
			log.Printf("failed to process assets: %v", err)
		}

		return
	}

	if *processor == transactionIngestor {
		transactionIngestor := service.NewTransactionIngestor(transactionRepo, accountRepo, assetRepo)
		err = transactionIngestor.ProcessTransactions(filePath, tabName, skipRows, accountID)
		if err != nil {
			log.Printf("failed to process transactions: %v", err)
		}

		return
	}

	//TODO: start gRPC service

	log.Println("Transaction service is started.")

}
