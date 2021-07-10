FROM golang:1.16-alpine
WORKDIR /src


COPY src/go.mod ./go.mod

RUN apk update \
  && apk add --no-cache git make musl-dev gcc \
  && go mod download
