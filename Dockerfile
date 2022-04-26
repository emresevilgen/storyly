FROM golang:1.17.5-alpine3.15 AS builder

ENV GO111MODULE=on

WORKDIR /build

RUN apk update && apk add git curl gcc musl-dev

RUN go get -u github.com/swaggo/swag/cmd/swag

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN swag init

RUN go build -o app

FROM alpine:3.11.0

WORKDIR /deployment

RUN apk add --update --no-cache ca-certificates && rm -rf /var/cache/apk/*

COPY --from=builder /build/app .

ENTRYPOINT ./app
