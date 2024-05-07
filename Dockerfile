# syntax=docker/dockerfile:1

FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY * ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-canteen-project

CMD ["/go-canteen-project"]