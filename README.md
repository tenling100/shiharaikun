# ShiharaiKun

## Overview

ShiharaiKun is a Go-based application designed to manage invoice data for future payment dates, automating bank transfers even if there's no balance on the due date. This allows users to delay cash outflows conveniently.

## Features

- REST API for managing invoices.
- Database interaction using GORM.
- Testing database operations with `sqlmock`.
- gRPC server for additional functionalities.
- Environment-based configuration management.

## Prerequisites

- Go 1.16+ installed (for local development)
- Docker and Docker Compose installed

## Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/tenling100/shiharaikun.git
   cd shiharaikun
   ```

2. **Set Up Environment Variables:**
   Create a `.env` file based on the `.env.example` provided and update it with your configuration.
   ```bash
   cp .env.example .env
   ```

3. **Install Go Modules:**
   ```bash
   go mod vendor
   ```

4. **Build the Application:**
   ```bash
   make build-server
   ```

## Database Migration

1. **Build the DB Migration Binary:**
   ```bash
   make build-migrate
   ```

2. **Run Database Migrations:**
   ```bash
   make migrate-db
   ```

### Reset DB
``` bash
make reset-db
```

## Generate Proto

To generate the protobuf files, follow these steps:

- **Install the Protocol Buffers Compiler (protoc):**
    - For macOS:
      ```bash
      make install-tools
      ```

- **Generate the Proto Files:**
    Run the following command to generate the protobuf files:
    ```bash
    make proto
    ```
    Replace `path/to/your/proto/file.proto` with the actual path to your proto file.

- **Format proto:**
  ```bash
  make proto-fmt
  ``` 


## Running the Project

### Using Docker Compose

1. **Build and Run Services with Docker Compose:**
   ```bash
   docker-compose up --build
   ```

   This command will build the necessary images and start the services defined in the `docker-compose.yml` file, including the database and the application server.

2. **Accessing the Application:**
   Once Docker Compose has started all services, the application will be accessible. Check the logs for the specific port it's running on or refer to the `docker-compose.yml` for configuration.

### Manually

1. **Run the Database:**
   Ensure your Mysql service is running and matches the configuration in the `.env` file.

2. **Run Database Migrations (if not already done):**
   ```bash
   make build-migrate
   make migrate-db
   ```

3. **Start the Application:**
   ```bash
   go run cmd/main.go
   ```

## Testing

### Testing with Docker Compose

1. **Start Services:**
   ```bash
   docker-compose up --build
   ```

2. **Run Database Migrations (if needed):**
   If the database schema needs to be updated or initialized, run:
   ```bash
   make migrate-db
   ```

3. **Run Test Script:**
   Navigate to the `test/` directory and run the desired test script:
   ```bash
   sh test/<script-name>.sh
   ```
   Replace `<script-name>` with the actual name of the script you want to run. Ensure that the test scripts are executable and properly configured with the necessary environment.

### Running Go Tests

Run the Go tests using the following command:

```bash
go test ./...
```

## Contributing

Feel free to fork this repository and submit pull requests.

## License

This project is licensed under the MIT License.
