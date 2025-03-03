package main

import (
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"log"

	"github.com/garcios/asset-trak-portfolio/asset-price-service/db"
	"github.com/garcios/asset-trak-portfolio/asset-price-service/handler"
	pb "github.com/garcios/asset-trak-portfolio/asset-price-service/proto"
	"github.com/garcios/asset-trak-portfolio/lib/mysql"
)

const (
	serviceName = "asset-price-service"
)

func main() {
	conn, err := mysql.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
	)

	srv.Init()

	assetPriceRepo := db.NewAssetPriceRepository(conn)

	h := handler.New(assetPriceRepo)

	err = pb.RegisterAssetPriceHandler(srv.Server(), h)
	if err != nil {
		log.Fatalf("failed to register currency handler: %v", err)
	}

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
	log.Println("Asset Price service is started.")
}
