include .env

tidy:
	go mod tidy

truncate:
	go run cmd/main.go -processor=truncate

ingest_assets:
	go run cmd/main.go -processor=assetIngestor

ingest_transactions: ingest_assets
	go run cmd/main.go -processor=transactionIngestor

run_server:
	go run cmd/main.go


.PHONY: tidy ingest_assets ingest_transactions run_server truncate

