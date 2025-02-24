include .env

init:
	go get -u google.golang.org/protobuf/proto
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest

tidy:
	go mod tidy

lint:
	golangci-lint run

run:
	DBUSER=${DB_USER} DBPASS=${DB_PASSWORD} go run cmd/main.go

proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/asset_price.proto

docker:
	podman build -t asset-price-service:1.0 .

.PHONY: init tidy run  proto docker lint

