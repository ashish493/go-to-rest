FROM golang:1.15-alpine AS build-env

RUN apk update && apk upgrade && \
    apk add --no-cache git

WORKDIR /  

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ./out/api 

FROM alpine:latest

RUN apk add ca-certificates

COPY --from=build /out/api /api

WORKDIR "/app"
EXPOSE 5000

CMD ["./api"]