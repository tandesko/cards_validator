FROM golang:1.21-alpine as buildbase

RUN apk add git build-base
ENV GO111MODULE=on

WORKDIR /go/src/github.com/tandesko/cards_validator

COPY . .
RUN go mod vendor

RUN GOOS=linux go build -o /usr/local/bin/cards_validator github.com/tandesko/cards_validator


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/cards_validator /usr/local/bin/cards_validator
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["cards_validator"]
