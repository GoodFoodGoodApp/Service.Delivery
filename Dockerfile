FROM golang:1.23-alpine

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY . /app/

RUN go mod download

# CMD ["air", "-c", ".air.toml"]
