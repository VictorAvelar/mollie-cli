FROM golang:alpine

ENV CGO_ENABLED=0

WORKDIR /testing

COPY go.* ./

RUN go mod download

COPY . .

RUN go test -v ./...