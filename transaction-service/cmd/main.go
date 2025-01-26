package main

import (
	"flag"
	"github.com/garcios/asset-track-portfolio/lib/mysql"
	"github.com/garcios/asset-track-portfolio/transactions/db"
	"github.com/garcios/asset-track-portfolio/transactions/service"
	"log"
)

const (
	assetIngestor       = "assetIngestor"
	transactionIngestor = "transactionIngestor"
	skipRows            = 3
	tabName             = "Combined"
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

	if *processor == assetIngestor {
		assetIngestor := service.NewAssetIngestor(assetRepo)

		err = assetIngestor.ProcessAssets(tabName, skipRows)
		if err != nil {
			log.Printf("failed to process assets: %v", err)
		}

		return
	}

	if *processor == transactionIngestor {
		transactionIngestor := service.NewTransactionIngestor(transactionRepo)
		err = transactionIngestor.ProcessTransactions(tabName, skipRows)
		if err != nil {
			log.Printf("failed to process transactions: %v", err)
		}

		return
	}

	//TODO: start gRPC service

	log.Println("Transaction service is started.")

}
