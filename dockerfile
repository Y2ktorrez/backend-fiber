FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

# Runtime image
FROM debian:bullseye
WORKDIR /root/

COPY --from=builder /app/main ./
COPY --from=builder /app/.env ./

EXPOSE 8000
CMD ["./main"]