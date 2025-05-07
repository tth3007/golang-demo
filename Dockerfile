# Use the official Go image as the base
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o myapp

# Use a lightweight image for the final executable
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp .

# Run the application
CMD ["./myapp"]