include .env

create_migration:
	migrate create -ext=sql -dir=. -seq init

migrate_up:
	migrate -path=. -database "${DB_APP}://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?multiStatements=true" -verbose up

migrate_down:
	migrate -path=. -database "${DB_APP}://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?multiStatements=true" -verbose down

uuidgen:
	uuidgen | tr '[:upper:]' '[:lower:]'

.PHONY: create_migration migrate_up migrate_down uuidgen

