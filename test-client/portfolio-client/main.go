package main

import (
	"context"
	"fmt"
	pb "github.com/garcios/asset-trak-portfolio/portfolio-service/proto"
	"go-micro.dev/v4"
)

const (
	ServiceName = "portfolio-service"
)

func main() {
	// Create a new service
	portfolioClient := micro.NewService(micro.Name("transaction-client"))
	portfolioClient.Init()

	transactionSrv := pb.NewTransactionService(ServiceName, portfolioClient.Client())

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
