package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4"
	"log"

	pb "github.com/garcios/asset-trak-portfolio/currency-service/proto"
)

const (
	ServiceName = "currency-service"
)

func main() {

	// Create a new service
	service := micro.NewService(micro.Name("currency-client"))
	service.Init()

	currencyService := pb.

	req := &pb.BalanceSummaryRequest{
		AccountId: "",
	}

	resp, err := transactionService.GetBalanceSummary(context.Background(), req)
	if err != nil {
		log.Fatalf("Get balance summary error: %v", err)
	}

	fmt.Println(resp)

}
