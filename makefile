## proto configuration
# Makefile for generating Go code from Protobuf files
PROTO_SRC_DIR := ./proto
PROTO_DST_DIR := ./api

# Define the protobuf compiler and Go plugin paths
PROTOC := protoc
PROTOC_GEN_GO := $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(shell go env GOPATH)/bin/protoc-gen-go-grpc
PROTOC_GEN_GRPC_GATEWAY := $(shell go env GOPATH)/bin/protoc-gen-grpc-gateway
PROTOC_GEN_OPENAPIV2 := $(shell go env GOPATH)/bin/protoc-gen-openapiv2
PROTOC_GEN_DOC := $(shell go env GOPATH)/bin/protoc-gen-doc
GOOGLE_APIS_PATH := ./third_party/googleapis

# List all the proto files to compile
PROTO_FILES := $(wildcard $(PROTO_SRC_DIR)/*.proto)

BUF := buf
PROTO_DIR = ./proto/

# Define variables for Docker
IMAGE_NAME = go-grpc-server
CONTAINER_NAME = go-grpc-server-container
DOCKER_PORT = 50051

# Define variables for Docker Compose
COMPOSE_FILE = docker-compose.yml

## Database configuration
# Variables for paths
MIGRATE_SCRIPT = cmd/migrate/main.go
MIGRATE_BIN = migrate
SCHEMA_FILE = db/schema.sql

# Define service name
DB_SERVICE_NAME = mysql

# Define Docker Compose file
DOCKER_COMPOSE_FILE = docker-compose.yml

# Define database credentials
DB_USER = testuser
DB_PASSWORD = testpassword
DB_NAME = testdb
DB_HOST = 127.0.0.1
DB_PORT = 3306

# Generate Go code from proto files
.PHONY: proto
proto: 
	@for proto in $(PROTO_FILES); do \
		echo "Generating Go code for $$proto"; \
		$(PROTOC) -I=$(PROTO_SRC_DIR) -I=$(GOOGLE_APIS_PATH) \
		          --go_out=$(PROTO_DST_DIR) --go_opt=paths=source_relative \
		          --go-grpc_out=$(PROTO_DST_DIR) --go-grpc_opt=paths=source_relative \
		          --grpc-gateway_out=$(PROTO_DST_DIR) --grpc-gateway_opt=paths=source_relative \
		          $$proto; \
	done

proto-fmt:
	@echo "Formatting proto files..."
	$(BUF) format -w $(PROTO_DIR)
	@echo "Proto files formatted."

proto-doc:
	@echo "Generating documentation for proto files..."
	@for proto in $(PROTO_FILES); do \
        echo "Generating documentation for $$proto"; \
        $(PROTOC) -I=$(PROTO_SRC_DIR) -I=$(GOOGLE_APIS_PATH) \
                  --doc_out=$(PROTO_DST_DIR) --doc_opt=markdown,README.md \
                  $$proto; \
    done
	@echo "Documentation generated."

# Install required tools
install-tools:
	@echo "Installing Protobuf tools"
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest


# Start the MySQL service using Docker Compose
start-db:
	@echo "Starting MySQL test database..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d $(DB_SERVICE_NAME)
	@echo "MySQL test database started."

renew-db:
	@echo "Renewing database: $(DB_NAME)"
	@mysql -u $(DB_USER) -p$(DB_PASS) -h $(DB_HOST) -e "DROP DATABASE IF EXISTS $(DB_NAME);"
	@mysql -u $(DB_USER) -p$(DB_PASS) -h $(DB_HOST) -e "CREATE DATABASE $(DB_NAME);"
	@echo "Database $(DB_NAME) has been renewed."

# Clean test environment
clean:
	@echo "Cleaning up test environment..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down -v
	@echo "Test environment cleaned."

# Run tests
.PHONY: test
test:
	# Run tests
	go test ./... -v

# build the binary
build-server:
	@echo "Building server binary..."
	go build -o server cmd/server/main.go
	@echo "Server binary built successfully."

# Target to build the migration binary
build-migrate:
	@echo "Building migration binary..."
	go build -o $(MIGRATE_BIN) $(MIGRATE_SCRIPT)
	@echo "Migration binary built successfully."

# Run the Go server
run: build
	@echo "Running the Go server..."
	./server

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME) .
	@echo "Docker image built: $(IMAGE_NAME)"

# Run Docker container
docker-run: docker-build
	@echo "Running Docker container..."
	docker run -d --name $(CONTAINER_NAME) -p $(DOCKER_PORT):$(DOCKER_PORT) $(IMAGE_NAME)
	@echo "Docker container running: $(CONTAINER_NAME)"

# Stop Docker container
docker-stop:
	@echo "Stopping Docker container..."
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true
	@echo "Docker container stopped and removed: $(CONTAINER_NAME)"

# Start services using Docker Compose
compose-up:
	@echo "Starting services with Docker Compose..."
	docker-compose -f $(COMPOSE_FILE) up -d --build
	@echo "Services started."

# Stop services using Docker Compose
compose-down:
	@echo "Stopping services with Docker Compose..."
	docker-compose -f $(COMPOSE_FILE) down
	@echo "Services stopped and removed."

# Target to run the migration binary
migrate-db:
	@echo "Running database migration..."
	./$(MIGRATE_BIN)
	@echo "Database migration completed."

# Target to dump the current database schema into schema.sql
dump-schema:
	@echo "Dumping database schema to $(SCHEMA_FILE)..."
	mysqldump -h $(DB_HOST) -P $(DB_PORT) -u$(DB_USER) -p$(DB_PASSWORD) \
	          --no-data --compact --skip-set-charset --skip-comments --no-tablespaces --skip-triggers --skip-routines $(DB_NAME) > $(SCHEMA_FILE)
	@echo "Database schema dumped to $(SCHEMA_FILE)."