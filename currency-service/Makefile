
init:
	go get -u google.golang.org/protobuf/proto
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest

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

docker:
	podman build -t currency-service:1.0 .

push:
	minikube image tag localhost/currency-service:1.0 localhost/currency-service:latest &&  minikube image push localhost/currency-service:latest

.PHONY: tidy ingest_currency_rates run_server truncate proto docker

