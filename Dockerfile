FROM golang:alpine

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY . /app/

RUN go mod download