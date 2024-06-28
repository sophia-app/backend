.PHONY: default run build test docs clean

# Variables
APP_NAME=storify

# Tasks
default: run-with-docs

run:
	@echo "Running $(APP_NAME)..."
	@go run ./cmd/$(APP_NAME)

run-with-docs:
	@make docs
	@make run

build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(APP_NAME) ./cmd/$(APP_NAME)

test:
	@echo "Running tests..."
	@go test -v ./...

docs:
	@echo "Generating docs..."
	@swag init -g ./cmd/$(APP_NAME)/main.go

clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME)
	@rm -rf ./docs
