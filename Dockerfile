#FROM golang:1.4.2

#RUN mkdir -p /go/src/coban/0.1/server/api
#COPY ./api /go/src/coban/0.1/server/api
#COPY ./configurations/environments.yml /go/src/coban/0.1/server/api/configurations/

#WORKDIR /go/src/coban/0.1/server/api

#RUN go get github.com/codegangsta/gin
#RUN go-wrapper download
#RUN go-wrapper install

#ENV DB_ENV development
#ENV PORT 3001

#EXPOSE 3000