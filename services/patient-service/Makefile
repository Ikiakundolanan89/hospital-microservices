
# Makefile
.PHONY: help build run test clean docker-build docker-run

# Variables
APP_NAME=patient-service
DOCKER_IMAGE=$(APP_NAME):latest
GO=go
GOFLAGS=-v

help: ## Display this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	$(GO) build $(GOFLAGS) -o bin/$(APP_NAME) ./cmd/main.go

run: ## Run the application
	$(GO) run ./cmd/main.go

test: ## Run tests
	$(GO) test $(GOFLAGS) ./...

clean: ## Clean build artifacts
	rm -rf bin/
	$(GO) clean

docker-build: ## Build Docker image
	docker build -t $(DOCKER_IMAGE) .

docker-run: ## Run Docker container
	docker run -p 3001:3001 --env-file .env $(DOCKER_IMAGE)

migrate-up: ## Run database migrations up
	migrate -path ./internal/database/migrations -database "sqlserver://sa:YourStrong@Passw0rd@localhost:1433?database=hospital_patient_db" up

migrate-down: ## Run database migrations down
	migrate -path ./internal/database/migrations -database "sqlserver://sa:YourStrong@Passw0rd@localhost:1433?database=hospital_patient_db" down

swagger: ## Generate Swagger documentation
	swag init -g ./cmd/main.go -o ./docs

dev: ## Run with hot reload (requires air)
	air