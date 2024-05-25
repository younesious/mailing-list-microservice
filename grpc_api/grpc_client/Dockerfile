FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o client ./grpc_api/grpc_client/grpc_client.go

# Start a new stage from scratch
FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=builder /app/client .

CMD ["./client"]
