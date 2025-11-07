# === Stage 1: Build stage ===
FROM golang:1.25.3-alpine AS builder

# Set build environment (default: development)
ARG NODE_ENV=development
ENV NODE_ENV=$NODE_ENV
ENV PATH="/go/bin:${PATH}"

# Set working directory
WORKDIR /app

# Copy go mod files dan download dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Install dependencies for Air (for development)
RUN if [ "$NODE_ENV" = "development" ]; then \
    go install github.com/air-verse/air@latest; \
    fi

# Build binary tergantung environment
# Kalau production → build dengan flags optimasi
RUN if [ "$NODE_ENV" = "production" ]; then \
    go build -o main -ldflags="-s -w" . ; \
    else \
    go build -o main . ; \
    fi


# === Stage 2: Runtime stage ===
FROM golang:1.25.3-alpine

WORKDIR /app

ENV PATH="/go/bin:${PATH}"

# Install Air
RUN go install github.com/air-verse/air@latest

# Copy semua source (buat Air bisa watch)
COPY . .

EXPOSE 3001

CMD ["air"]
