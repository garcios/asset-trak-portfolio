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

run:
	DBUSER=root DBPASS=Pass123 go run cmd/main.go

proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/asset_price.proto

docker:
	podman build -t asset-price-service:1.0 .

.PHONY: init tidy ingest_asset_prices run_server truncate proto docker

