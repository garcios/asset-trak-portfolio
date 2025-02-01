package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4"
	"log"

	pb "github.com/garcios/asset-trak-portfolio/transaction-service/proto"
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
		AccountId: "",
	}

	resp, err := transactionService.GetBalanceSummary(context.Background(), req)
	if err != nil {
		log.Fatalf("Get balance summary error: %v", err)
	}

	fmt.Println(resp)

}
