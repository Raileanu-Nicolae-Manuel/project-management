.PHONY: migrate-up migrate-down db-create dev test build run sqlc-generate

# Include environment variables from .env file
include .env
export

migrate-up:
	migrate -database "mysql://${DATABASE_URL}" -path ./internal/db/migrations up

migrate-down:
	migrate -database "mysql://${DATABASE_URL}" -path ./internal/db/migrations down 

db-create: 
	docker-compose build; \
	docker-compose up -d

dev:
	air main.go

test:
	go test -v ./...

build:
	go build -o backend main.go

run:
	./backend

sqlc-generate:
	sqlc generate



