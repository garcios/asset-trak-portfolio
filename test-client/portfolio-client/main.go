package main

import (
	"context"
	"fmt"
	lib "github.com/garcios/asset-trak-portfolio/lib/retryable"
	pb "github.com/garcios/asset-trak-portfolio/portfolio-service/proto"
	"time"
)

const (
	ServiceName = "portfolio-service"
)

func main() {
	// Create a new client service
	portfolioClient := lib.CreateRetryableClient(
		"portfolio-client",
		lib.WithMaxRetries(3),             // optional
		lib.WithRetryDelay(2*time.Second), // optional
	)
	portfolioClient.Init()

	transactionSrv := pb.NewPortfolioService(ServiceName, portfolioClient.Client())

	req := &pb.HoldingsRequest{
		AccountId: "eb08df3c-958d-4ae8-b3ae-41ec04418786",
	}

	resp, err := transactionSrv.GetHoldings(context.Background(), req)
	if err != nil {
		fmt.Printf("Get balance summary error: %v", err)
		return
	}

	for _, item := range resp.Investments {
		fmt.Println(item)
	}

}
