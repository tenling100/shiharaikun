## proto configuration
# Makefile for generating Go code from Protobuf files
PROTO_SRC_DIR := proto
PROTO_DST_DIR := ./api

# Define the protobuf compiler and Go plugin paths
PROTOC := protoc
PROTOC_GEN_GO := $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(shell go env GOPATH)/bin/protoc-gen-go-grpc

# List all the proto files to compile
PROTO_FILES := $(wildcard $(PROTO_SRC_DIR)/*.proto)


## Database configuration
# Variables for paths
MIGRATE_SCRIPT = cmd/migrate/main.go
MIGRATE_BIN = migrate
SCHEMA_FILE = db/schema.sql

# Define service name
DB_SERVICE_NAME = test-mysql

# Define Docker Compose file
DOCKER_COMPOSE_FILE = docker-compose.yml

# Define database credentials
DB_USER = testuser
DB_PASSWORD = testpassword
DB_NAME = testdb
DB_HOST = 127.0.0.1
DB_PORT = 3306

# Default target
all: generate

# Generate Go code from proto files
generate: $(PROTO_FILES)
	@for proto in $(PROTO_FILES); do \
		echo "Generating Go code for $$proto"; \
		$(PROTOC) --plugin=protoc-gen-go=$(PROTOC_GEN_GO) --plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
				  --go_out=module=github.com/tenling100/shiharaikun/api:$(PROTO_DST_DIR) \
				  --go-grpc_out=module=github.com/tenling100/shiharaikun/api:$(PROTO_DST_DIR) \
		          $$proto; \
	done

# Install required tools
install-tools:
	@echo "Installing Protobuf tools"
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


# Start the MySQL service using Docker Compose
start-db:
	@echo "Starting MySQL test database..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d $(DB_SERVICE_NAME)
	@echo "MySQL test database started."

# Stop and remove the MySQL service
stop-db:
	@echo "Stopping MySQL test database..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down
	@echo "MySQL test database stopped and removed."

# Initialize the test database schema
init-db:
	@echo "Initializing test database schema..."
	# Run SQL commands to initialize schema
	# This command uses mysql client to execute a schema.sql script
	docker-compose -f $(DOCKER_COMPOSE_FILE) exec $(DB_SERVICE_NAME) \
	bash -c "mysql -u $(DB_USER) -p$(DB_PASSWORD) $(DB_NAME) < /path/to/schema.sql"
	@echo "Test database schema initialized."

# Clean test environment
clean:
	@echo "Cleaning up test environment..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down -v
	@echo "Test environment cleaned."

# Run tests
test:
	@echo "Running tests..."
	# Ensure the DB is up
	make start-db
	# Run tests
	go test ./... -v
	# Stop the DB after tests
	make stop-db
	@echo "Tests completed."

# Target to build the migration binary
build-migrate:
	@echo "Building migration binary..."
	go build -o $(MIGRATE_BIN) $(MIGRATE_SCRIPT)
	@echo "Migration binary built successfully."

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