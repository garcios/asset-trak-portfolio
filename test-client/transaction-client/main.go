package main

import (
	"context"
	"fmt"
	pb "github.com/garcios/asset-trak-portfolio/transaction-service/proto"
	"go-micro.dev/v4"
)

const (
	ServiceName = "transaction-service"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("transaction-client"))
	service.Init()

	transactionService := pb.NewTransactionService(ServiceName, service.Client())

	req := &pb.BalanceSummaryRequest{
		AccountId: "eb08df3c-958d-4ae8-b3ae-41ec04418786",
	}

	resp, err := transactionService.GetBalanceSummary(context.Background(), req)
	if err != nil {
		fmt.Printf("Get balance summary error: %v", err)
		return
	}

	fmt.Printf("Account ID: %v\n", resp.AccountId)
	fmt.Printf("Total Value: %v\n", resp.TotalValue)

	for _, item := range resp.BalanceItems {
		fmt.Println(item)
	}

}
