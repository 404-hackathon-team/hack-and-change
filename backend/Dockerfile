# ---- Build stage ----
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/server ./cmd

# ---- Run stage ----
FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/server .

COPY ./cmd/migrations ./migrations

# папки под данные
RUN mkdir -p /app/uploads /app/data /app/logs


CMD ["./server"]