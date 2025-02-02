include .env

tidy:
	go mod tidy

truncate:
	go run cmd/main.go -processor=truncate

ingest_currency_rates:
	go run cmd/main.go -processor=currencyRateIngestor

run_server:
	go run cmd/main.go

proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/currency.proto

.PHONY: tidy ingest_currency_rates run_server truncate proto

