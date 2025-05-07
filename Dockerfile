# syntax=docker/dockerfile:1
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /bin/app .

CMD ["./app"]