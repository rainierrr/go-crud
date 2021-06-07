FROM golang:1.16-alpine
WORKDIR /go/src/


COPY go/src/go.mod ./go.mod
COPY go/src/go.sum ./go.sum

RUN apk update \
  && apk add --no-cache git make \
  && go mod download
