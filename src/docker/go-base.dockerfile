# podman build -f docker/go-base.dockerfile -t go-base:latest .

FROM golang:1.23 AS builder

WORKDIR /src
COPY go.mod go.sum .

RUN go mod download

COPY kitex_gen/ ./kitex_gen/
