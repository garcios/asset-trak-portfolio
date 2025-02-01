include .env

tidy:
	go mod tidy

truncate:
	go run cmd/main.go -processor=truncate

ingest_currency_rates:
	go run cmd/main.go -processor=currencyRateIngestor

run_server:
	go run cmd/main.go


.PHONY: tidy ingest_currency_rates run_server truncate

