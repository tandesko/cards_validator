FROM golang:1.21-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/tandesko/cards_validator
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/cards_validator github.com/tandesko/cards_validator


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/cards_validator /usr/local/bin/cards_validator
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["cards_validator"]
