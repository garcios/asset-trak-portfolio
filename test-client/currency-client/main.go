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

	/*	response, err := currencySrv.GetExchangeRate(
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

		log.Println(response)*/

	res, err := currencySrv.GetHistoricalExchangeRates(context.Background(), &pb.GetHistoricalExchangeRatesRequest{
		FromCurrency: "USD",
		ToCurrency:   "AUD",
		StartDate:    "2020-07-01",
		EndDate:      "2025-01-30",
	})

	if err != nil {
		log.Printf("Error calling GetHistoricalExchangeRates: %v", err)
		return
	}

	for _, rec := range res.HistoricalRates {
		log.Println(rec)
	}

}
