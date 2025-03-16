# ------------------------------
# 🛠️ Stage 1: Build the Go binary
# ------------------------------
FROM golang:1.23-alpine AS builder

# Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    PATH="/go/bin:${PATH}"

# Create app directory
WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Install tools
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy go mod files first for cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source
COPY . .

# Generate Swagger docs
RUN swag init -g app/main.go

# Build the app
RUN go build -o task-manager ./app

# ------------------------------
# 🚀 Stage 2: Final runtime image
# ------------------------------
FROM alpine:latest

# Create non-root user
RUN adduser -D appuser

# Set work directory
WORKDIR /app

# Copy built binary and required assets
COPY --from=builder /app/task-manager .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /go/bin/swag /usr/local/bin/swag

# Set user
USER appuser

# Expose API port
EXPOSE 8080

# CMD
CMD ["./task-manager"]
