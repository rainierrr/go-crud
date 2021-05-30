FROM golang:1.16-alpine
WORKDIR /code


COPY code/go.mod ./go.mod
COPY code/go.sum ./go.sum

RUN apk update \
  && apk add --no-cache git
RUN go mod download

#RUN go install github.com/gin-gonic/gin@v1.7.4
  #&& go mod tidy
  #&& go get github.com/jinzhu/gorm \
  #&&
