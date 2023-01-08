FROM golang:1.19-alpine

WORKDIR /go/src
COPY ./app/src .

RUN apk upgrade --update && apk --no-cache add git

RUN  go install github.com/spf13/cobra-cli@latest
