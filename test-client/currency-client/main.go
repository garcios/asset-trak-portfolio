package main

import (
	"context"
	"fmt"
	pb "github.com/garcios/asset-trak-portfolio/currency-service/proto"
	"go-micro.dev/v4"
)

const (
	ServiceName = "currency-service"
)

func main() {
	// Create a new service
	currencyClient := micro.NewService(micro.Name("currency-client"))
	currencyClient.Init()

	currencySrv := pb.NewCurrencyService(ServiceName, currencyClient.Client())

	response, err := currencySrv.GetExchangeRate(
		context.Background(),
		&pb.GetExchangeRateRequest{
			FromCurrency: "USD",
			ToCurrency:   "AUD",
			TradeDate:    "2025-01-30",
		},
	)
	if err != nil {
		fmt.Printf("Error calling GetExchangeRate: %v", err)
		return
	}

	fmt.Println(response)
}
