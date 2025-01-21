# Build
FROM golang:1.23.2-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Install necessary dependencies
RUN apk add --no-cache gcc musl-dev

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application as a statically linked binary
RUN CGO_ENABLED=0 go build -o main .

# Minimal runtime image
FROM alpine:3.21

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port that the application will listen on
EXPOSE 8080

# Run the Go application
CMD ["./main"]
