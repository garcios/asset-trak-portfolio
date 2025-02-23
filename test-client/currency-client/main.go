package main

import (
	"context"
	pb "github.com/garcios/asset-trak-portfolio/currency-service/proto"
	lib "github.com/garcios/asset-trak-portfolio/lib/retryable"
	"log"
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
		log.Printf("Error calling GetExchangeRate: %v", err)
		return
	}

	log.Println(response)
}
