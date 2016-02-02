FROM golang:1.4.2

RUN mkdir -p /go/src/coban/api/0.1/server/api
COPY . /go/src/coban/api/0.1/server/api
WORKDIR /go/src/coban/api/0.1/server/api

ENV GOPATH /go

#RUN go get github.com/codegangsta/gin
RUN go-wrapper download
RUN go-wrapper install