package main

import (
	"context"
	pb "github.com/garcios/asset-trak-portfolio/asset-price-service/proto"
	lib "github.com/garcios/asset-trak-portfolio/lib/retryable"
	"log"
	"time"
)

const (
	ServiceName = "asset-price-service"
)

func main() {
	// Create a new client service
	apClient := lib.CreateRetryableClient(
		"portfolio-client",
		lib.WithMaxRetries(3),             // optional
		lib.WithRetryDelay(2*time.Second), // optional
	)
	apClient.Init()

	apSrv := pb.NewAssetPriceService(ServiceName, apClient.Client())

	//req := &pb.GetAssetPriceRequest{
	//	AssetId:   "2264c1ab-0adf-4349-afb4-b1694e7f97c1",
	//	TradeDate: "2025-01-30",
	//}

	//resp, err := apSrv.GetAssetPrice(context.Background(), req)
	//if err != nil {
	//	log.Printf("Get asset price error: %v", err)
	//	return
	///}

	//log.Println(resp)

	req2 := &pb.GetAssetPriceHistoryRequest{
		AssetId:   "2264c1ab-0adf-4349-afb4-b1694e7f97c1",
		StartDate: "2020-07-01",
		EndDate:   "2025-01-30",
	}

	res2, err := apSrv.GetAssetPriceHistory(context.Background(), req2)
	if err != nil {
		log.Printf("Get asset prices error: %v", err)
		return
	}

	for _, p := range res2.GetPrices() {
		log.Println(p)
	}

	log.Println(res2)

}
