FROM library/golang:1.16.2-alpine3.13

RUN apk add make gcc libc-dev bash

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest