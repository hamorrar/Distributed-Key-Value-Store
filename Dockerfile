# syntax=docker/dockerfile:1

FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o kvs ./src

EXPOSE 8090

CMD ./kvs