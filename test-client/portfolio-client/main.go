package main

import (
	"context"
	lib "github.com/garcios/asset-trak-portfolio/lib/retryable"
	pb "github.com/garcios/asset-trak-portfolio/portfolio-service/proto"
	"log"
	"time"
)

const (
	ServiceName = "portfolio-service"
)

func main() {
	// Create a new client service
	portfolioClient := lib.CreateRetryableClient(
		"portfolio-client",
		lib.WithMaxRetries(3),              // optional
		lib.WithRetryDelay(30*time.Second), // optional
	)
	portfolioClient.Init()

	transactionSrv := pb.NewPortfolioService(ServiceName, portfolioClient.Client())

	//req := &pb.HoldingsRequest{
	//	AccountId: "eb08df3c-958d-4ae8-b3ae-41ec04418786",
	//}

	//resp, err := transactionSrv.GetHoldings(context.Background(), req)
	//if err != nil {
	//	log.Printf("Get balance summary error: %v", err)
	//	return
	//}

	//for _, item := range resp.Investments {
	//	log.Println(item)
	//}

	preq := &pb.PerformanceHistoryRequest{
		AccountId: "eb08df3c-958d-4ae8-b3ae-41ec04418786",
		StartDate: "2024-09-01",
		EndDate:   "2025-03-03",
	}

	presp, err := transactionSrv.GetPerformanceHistory(context.Background(), preq)
	if err != nil {
		log.Printf("Get performance history error: %v", err)
		return
	}

	for _, r := range presp.Records {
		log.Println(r)
	}

}
