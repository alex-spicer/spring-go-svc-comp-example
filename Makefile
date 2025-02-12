.PHONY: all build clean run test docker-build docker-run

# Build all services
build: build-java build-go

build-java:
	@echo "Building Java service..."
	cd product && ./gradlew build -x test

build-go:
	@echo "Building Go service..."
	cd product-go && go build ./cmd/api

# Run services locally
run-java:
	@echo "Running Java service..."
	cd product && ./gradlew bootRun

run-go:
	@echo "Running Go service..."
	cd product-go && go run cmd/api/main.go

# Docker commands
docker-build:
	docker-compose build

docker-run:
	docker-compose up

docker-stop:
	docker-compose down