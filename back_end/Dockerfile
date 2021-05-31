FROM golang:1.16-alpine
WORKDIR /code


COPY code/go.mod ./go.mod
COPY code/go.sum ./go.sum

RUN apk update \
  && apk add --no-cache git make \
  && go mod download
