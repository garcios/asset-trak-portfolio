init:
	go get -u google.golang.org/protobuf/proto
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest


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

proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/transaction.proto

docker:
	podman build -t transaction-service:1.0 .

.PHONY: init tidy ingest_assets ingest_transactions run_server truncate proto docker

