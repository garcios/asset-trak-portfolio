package main

import (
	"go-micro.dev/v4"

	pb "github.com/garcios/asset-trak-portfolio/transaction-service/proto"
)

const (
	ServiceName = "transaction-service"
)

func main() {

	// Create a new service
	service := micro.NewService(micro.Name("transaction-client"))
	service.Init()

}
