include .env

run:
	go run main.go

mod-vendor:
	go mod vendor

golangci-lint:
	@golangci-lint run

gosec:
	@gosec -quiet ./...

validate: golangci-lint gosec

docker:
	docker-compose build
	docker-compose up

migrate-create:
	@goose -dir=migrations create "$(name)" sql

migrate-up:
	@goose -dir=migrations postgres "host=${POSTGRES_HOST} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" up

migrate-down:
	@goose -dir=migrations postgres "host=${POSTGRES_HOST} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" down

swagger: swag redocly swag-fmt

swag:
	swag init -g main.go

redocly:
	redocly build-docs docs/swagger.yaml -o docs/index.html

install-swagger:
	go install github.com/swaggo/swag/cmd/swag@latest

get-swagdeps:
	go get github.com/swaggo/swag

swag-fmt:
	swag fmt
