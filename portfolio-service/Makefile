include .env

init:
	go get -u google.golang.org/protobuf/proto
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest


tidy:
	go mod tidy

truncate:
	DBUSER=${DB_USER} DBPASS=${DB_PASSWORD} go run cmd/main.go -processor=truncate

ingest-trades:
	DBUSER=${DB_USER} DBPASS=${DB_PASSWORD} go run cmd/main.go -processor=tradesIngestor -accountID="eb08df3c-958d-4ae8-b3ae-41ec04418786"

ingest-dividends:
	DBUSER=${DB_USER} DBPASS=${DB_PASSWORD} go run cmd/main.go -processor=dividendsIngestor -accountID="eb08df3c-958d-4ae8-b3ae-41ec04418786"

run:
	DBUSER=${DB_USER} DBPASS=${DB_PASSWORD} go run cmd/main.go

proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/portfolio.proto

docker:
	podman build -t portfolio-service:1.0 .

.PHONY: init tidy ingest_assets ingest-trades ingest-dividends run truncate proto docker

