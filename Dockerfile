FROM golang:1.18-alpine

ENV CGO_ENABLED=0

RUN apk add wget

RUN go install github.com/VictorAvelar/mollie-cli/cmd/mollie@latest

RUN wget https://raw.githubusercontent.com/VictorAvelar/mollie-cli/master/.mollie.yaml

ENTRYPOINT ["mollie"]

CMD ["-h"]