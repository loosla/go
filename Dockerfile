# go1.18.1 linux/amd64
FROM golang@sha256:f1e97d64a50f4c2b4fa61211f5206e636a54f992a047d192d6d068fbcd1946c3 as builder


WORKDIR /app

COPY go.mod ./
# COPY go.sum ./
RUN go mod download

COPY * ./

RUN go build -o /main

ENTRYPOINT ["/main"]
