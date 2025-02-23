package main

import (
	"fmt"
	"github.com/garcios/asset-trak-portfolio/lib/mysql"
	"log"
)

func main() {
	_, err := mysql.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("Asset Price service is being started...")

	//TODO: start gRPC service

	log.Println("Asset Price service is started.")
}
