# syntax=docker/dockerfile:1

FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/api
RUN go build -o /app/app .

# --- Release image ---
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

# Копируем .env, если он есть
COPY .env .env

ENV GIN_MODE=release

EXPOSE 8080

CMD ["./app"] 