include .env

init:
	go get -u google.golang.org/protobuf/proto
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest

tidy:
	go mod tidy

truncate-assets:
	DBUSER=${DB_USER} DBPASS=${DB_PASSWORD} go run cmd/main.go -processor=truncateAsset

truncate-asset-prices:
	DBUSER=${DB_USER} DBPASS=${DB_PASSWORD} go run cmd/main.go -processor=truncateAssetPrice

truncate-currency-rates:
	DBUSER=${DB_USER} DBPASS=${DB_PASSWORD} go run cmd/main.go -processor=truncateCurrencyRates

ingest-asset:
	DBUSER=${DB_USER} DBPASS=${DB_PASSWORD} go run cmd/main.go -processor=ingestAsset

ingest-asset-prices:
	DBUSER=${DB_USER} DBPASS=${DB_PASSWORD} go run cmd/main.go -processor=ingestAssetPrice

ingest-currency-rates:
	DBUSER=${DB_USER} DBPASS=${DB_PASSWORD} go run cmd/main.go -processor=ingestCurrencyRates

run:
	DBUSER=${DB_USER} DBPASS=${DB_PASSWORD} go run cmd/main.go

proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/asset_price.proto

docker:
	podman build -t asset-price-service:1.0 .

.PHONY: init tidy  run  proto docker truncate-assets truncate-asset-prices ingest-asset ingest-asset-prices ingest-currency-rates

