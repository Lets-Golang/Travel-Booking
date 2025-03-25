FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/user-service .
COPY --from=builder /app/.env .
COPY --from=builder /app/migrations ./migrations

CMD ["./user-service"]