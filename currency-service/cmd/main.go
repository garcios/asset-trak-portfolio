package main

import (
	"github.com/garcios/asset-trak-portfolio/currency-service/db"
	"github.com/garcios/asset-trak-portfolio/currency-service/handler"
	pb "github.com/garcios/asset-trak-portfolio/currency-service/proto"
	"github.com/garcios/asset-trak-portfolio/lib/mysql"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"log"
)

const (
	serviceName = "currency-service"
)

func main() {
	conn, err := mysql.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	currencyRepo := db.NewCurrencyRepository(conn)

	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
	)

	srv.Init()

	h := handler.New(currencyRepo)

	err = pb.RegisterCurrencyHandler(srv.Server(), h)
	if err != nil {
		log.Fatalf("failed to register currency handler: %v", err)
	}

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}

	log.Println("Currency service is started.")
}
