package main

import (
	"context"
	"fmt"
	pb "github.com/garcios/asset-trak-portfolio/currency-service/proto"
	lib "github.com/garcios/asset-trak-portfolio/lib/retryable"
)

const (
	ServiceName = "currency-service"
)

func main() {
	// Create a new client service
	currencyClient := lib.CreateRetryableClient(
		"currency-client",
	)

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
