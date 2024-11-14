# Stage 1: Build the application
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy dependency files and download modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the application from the correct directory
RUN CGO_ENABLED=0 GOOS=linux go build -o dnq-backend ./cmd

# Stage 2: Create the final lightweight image
FROM alpine:latest

# Install SSL certificates for HTTPS communication
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/dnq-backend .

# Expose the application port
EXPOSE 8080

CMD ["./dnq-backend"]
