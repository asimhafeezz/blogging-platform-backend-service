FROM golang:1.24-alpine

WORKDIR /app

# Install necessary tools
RUN apk add --no-cache git

# Copy go mod files and download deps first (for cache)
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code (but will be overridden by volume in dev)
COPY . .

# Default command — just run main.go
CMD ["go", "run", "main.go"]