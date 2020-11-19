.PHONY: build
build: go-build ## Build Go services

.PHONY: clean
clean: go-clean ## Clean build cache and dependencies

.PHONY: gen-mocks
gen-mocks: go-gen-mocks ## Generate mocks, requires mockery to be installed (https://github.com/vektra/mockery)

.PHONY: unit-test
unit-test: go-unit-test ## Run unit tests

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

go-build:
	@echo "Building Go services..."
	@rm -rf build
	@mkdir build
	go build -o build -v ./...
	@echo "Go services available at ./build"

go-clean: go-clean-cache go-clean-deps

go-clean-cache:
	@echo "Cleaning build cache..."
	go clean -cache

go-clean-test-cache:
	@echo "Cleaning test cache..."
	go clean -testcache

go-clean-deps:
	@echo "Cleaning dependencies..."
	go mod tidy

go-deps:
	@echo "Installing dependencies..."
	go mod download

go-unit-test:
	@echo "Running unit tests..."
	go test -v

go-gen-mocks:
	@echo "Generating mocks..."
	# requires mockery tool https://github.com/vektra/mockery
	mockery --name=HTTPClient --output=mocks

services-build:
	@echo "Building services..."
	docker-compose -f docker-compose.yml build --force-rm

.PHONY: services-up
services-up:
	@echo "Starting all services..."
	docker-compose -f docker-compose.yml up --force-recreate --detach

.PHONY: services-down
services-down:
	@echo "Stopping all services..."
	docker-compose -f docker-compose.yml down

.DEFAULT_GOAL := help