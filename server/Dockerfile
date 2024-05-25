FROM golang:1.22-alpine as builder

RUN apk add --no-cache gcc musl-dev
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o server ./server/server.go

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

# Expose 8080 for HTTP and 8081 for gRPC 
EXPOSE 8080
EXPOSE 8081


CMD ["./server"]
