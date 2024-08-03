# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR ~/go/src/github.com/exhibit-io/tracker

COPY . .

RUN go build -o /tracker

EXPOSE 8080

CMD [ "/tracker" ]