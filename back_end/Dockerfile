FROM golang:1.16-alpine
WORKDIR /src


COPY src/go.mod ./go.mod
COPY src/go.sum ./go.sum

RUN apk update \
  && apk add --no-cache git make \
  && go mod download
