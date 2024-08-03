# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY *.go ./
COPY tracker ./tracker

RUN go mod tidy && go build -o /tracker

EXPOSE 8080

CMD [ "/tracker" ]