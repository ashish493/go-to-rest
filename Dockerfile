FROM golang:1.15-alpine AS build-env

WORKDIR /
RUN go build -o /go/bin/
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

EXPOSE 8080