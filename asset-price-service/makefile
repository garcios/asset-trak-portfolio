include .env

tidy:
	go mod tidy

truncate:
	go run cmd/main.go -processor=truncate

ingest_asset_prices:
	go run cmd/main.go -processor=assetPriceIngestor

run_server:
	go run cmd/main.go


.PHONY: tidy ingest_asset_prices run_server truncate

