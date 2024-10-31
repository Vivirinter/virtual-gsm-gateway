FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o virtual-gsm-gateway ./cmd/main

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/virtual-gsm-gateway .
EXPOSE 8080
CMD ["./virtual-gsm-gateway"]