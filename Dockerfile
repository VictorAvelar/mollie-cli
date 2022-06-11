FROM golang:1.18-alpine

ENV CGO_ENABLED=0

WORKDIR /testing

COPY go.* ./

RUN go mod download

COPY . .

ENTRYPOINT [ "go", "test", "./...", "-v"]