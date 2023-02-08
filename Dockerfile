# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY src/*.go ./

RUN go build -o /kvs

EXPOSE 8090

CMD /kvs