.PHONY: build run clean test docker-build docker-run help

# Variables
BINARY_NAME=workprofileapp
GO=go
DOCKER_IMAGE_NAME=workprofileapp
DOCKER_IMAGE_TAG=latest

# Default target
help:
	@echo "Available targets:"
	@echo "  make build         - Build the application"
	@echo "  make run           - Run the application locally"
	@echo "  make clean         - Remove build artifacts"
	@echo "  make test          - Run tests"
	@echo "  make docker-build  - Build Docker image"
	@echo "  make docker-run    - Run Docker container"
	@echo "  make help          - Show this help message"

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	CGO_ENABLED=0 $(GO) build -o $(BINARY_NAME) .
	@echo "Build complete: $(BINARY_NAME)"

# Run the application locally
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	$(GO) test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	$(GO) clean
	rm -f $(BINARY_NAME)
	@echo "Clean complete"

# Build Docker image
docker-build:
	@echo "Building Docker image: $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)..."
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) .
	@echo "Docker image built successfully"

# Run Docker container
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 --name $(DOCKER_IMAGE_NAME) $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

# Format code
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# Run linter (requires golangci-lint to be installed)
lint:
	@echo "Running linter..."
	golangci-lint run ./...
