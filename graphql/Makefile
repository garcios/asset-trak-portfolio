# variables
GEN_DIR := generated

clean:
	rm -rf ${GEN_DIR}

tidy:
	go mod tidy

lint:
	golangci-lint run

generate: clean
	go run github.com/99designs/gqlgen generate

run:
	go run server.go

.PHONY: clean generate  run tidy lint
