# Stage 1: Build the Go binary
FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main -v ./cmd/temperature-error-logs



# Stage 2: Run the Go binary
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
