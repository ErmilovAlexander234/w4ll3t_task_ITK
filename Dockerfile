FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o w4ll3t ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/w4ll3t .
COPY config.env .
EXPOSE 8080
CMD ["./w4ll3t"]