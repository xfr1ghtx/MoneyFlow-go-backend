# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o app .

# --- Release image ---
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

# Для godotenv (если .env нужен внутри контейнера)
COPY ../../.env .env

ENV GIN_MODE=release

EXPOSE 8080

CMD ["./app"] 