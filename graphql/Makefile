# variables
GEN_DIR := generated

clean:
	rm -rf ${GEN_DIR}

generate: clean
	go run github.com/99designs/gqlgen generate

run:
	go run server.go

.PHONY: clean generate  run
