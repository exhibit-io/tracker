# syntax=docker/dockerfile:1

FROM golang:1.22.5-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o /tracker

EXPOSE 8080

CMD [ "/tracker" ]