# syntax=docker/dockerfile:1

FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/api
RUN go build -o /app/app .

# Устанавливаем goose для миграций
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# --- Release image ---
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /go/bin/goose /usr/local/bin/goose

# Копируем .env, если он есть
COPY .env .env

ENV GIN_MODE=release

EXPOSE 8080

CMD ["./app"] 