FROM golang:1.26.2-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/user-cache-service ./cmd/app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user-cache-service .

EXPOSE 8080

CMD ["./user-cache-service"]