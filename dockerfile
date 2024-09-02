# Use the official Golang image as a base
FROM golang:1.23-alpine

# Install make utility
RUN apk add --no-cache make

# Set the Current Working Directory inside the container
WORKDIR /shiharaikun

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN make build-server
RUN make build-migrate

# Expose port 8080 to the outside world
EXPOSE 8080
EXPOSE 50051

# Copy the wait-for-mysql script into the container
COPY wait-for-mysql.sh /wait-for-mysql.sh
COPY entrypoint.sh /entrypoint.sh

# Make the scripts executable
RUN chmod +x /wait-for-mysql.sh /entrypoint.sh migrate

# Command to run the entrypoint script
ENTRYPOINT ["/entrypoint.sh"]



# Command to run the executable
CMD ["./server"]