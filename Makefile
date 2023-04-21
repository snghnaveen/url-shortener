# Set environment variables
ENVIRONMENT ?= testing
APP_PORT ?= 8081


tidy:
	go mod tidy

# Build the application
build:
	ENVIRONMENT=$(ENVIRONMENT) APP_PORT=$(APP_PORT) go build -o app .

# Run the application
run: tidy build
	ENVIRONMENT=$(ENVIRONMENT) APP_PORT=$(APP_PORT) ./app

# Clean up the build artifacts
clean:
	rm -f app

# Run tests
test:
	ENVIRONMENT=$(ENVIRONMENT) go test ./...

# Show help
help:
	@echo "make build - Build the application"
	@echo "make run - Run the application"
	@echo "make clean - Clean up the build artifacts"
	@echo "make test - Run tests"
