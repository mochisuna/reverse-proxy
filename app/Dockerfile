FROM golang:alpine as local
LABEL maintainer="mochisuna"

WORKDIR /app
COPY . .
RUN apk update && apk add --no-cache git make && \
    go install github.com/cosmtrek/air@latest && \
    make build
