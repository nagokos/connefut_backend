FROM golang:1.17

RUN apt-get update && apt-get install git

RUN mkdir /go/src/app
WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@latest

ADD . /go/src/app/s

CMD [ "air" ]