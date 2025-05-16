FROM golang:1.24-alpine

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application from the cmd/api directory
RUN go build -o main ./cmd/api

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"] 