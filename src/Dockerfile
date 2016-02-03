FROM golang:1.4.2

RUN mkdir -p /go/src/coban/api/
COPY . /go/src/coban/api
RUN mv ./databases/migrations /go/src/coban/migrations

ENV GOPATH /go

#
# Build the migrations
#

WORKDIR /go/src/coban/migrations
RUN go-wrapper download
RUN go-wrapper install

#
# Build the API
#

WORKDIR /go/src/coban/api

#RUN go get github.com/codegangsta/gin
RUN go-wrapper download
RUN go-wrapper install