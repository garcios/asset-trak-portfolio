init:
	go get -u google.golang.org/protobuf/proto
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest

tidy:
	go mod tidy

truncate:
	DBUSER=root DBPASS=Pass123 go run cmd/main.go -processor=truncate

ingest_asset_prices:
	DBUSER=root DBPASS=Pass123 go run cmd/main.go -processor=assetPriceIngestor

run_server:
	go run cmd/main.go

docker:
	podman build -t asset-price-service:1.0 .

.PHONY: init tidy ingest_asset_prices run_server truncate docker

