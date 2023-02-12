# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

COPY *.go ./
COPY .apiConfig ./

RUN go build -o /weather-app

EXPOSE 8080

CMD [ "/weather-app" ]