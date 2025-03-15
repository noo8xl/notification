# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS builder

WORKDIR /notificator/

COPY go.mod go.sum ./
COPY . .

RUN go mod download
RUN go mod tidy
RUN GOFLAGS=-mod=mod GOOS=linux GOARCH=amd64 go build -o bin/notificator main.go

EXPOSE 7493

CMD ["./bin/notificator"] 

