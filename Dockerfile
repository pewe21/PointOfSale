# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder


# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod tidy

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o myapp ./cmd

# Stage 2: Create a smaller image
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./myapp"]
