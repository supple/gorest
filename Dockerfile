#FROM debian:jessie
FROM golang:1.6.2-alpine

ENV APP_PKG_DIR /go/src/github.com/supple/mtest
WORKDIR $APP_PKG_DIR

RUN mkdir /data

ADD ./ $APP_PKG_DIR
RUN  go build -v -o mtest .

ENTRYPOINT ["/go/src/github.com/supple/mtest/mtest"]

EXPOSE 8000
VOLUME ["/data"]