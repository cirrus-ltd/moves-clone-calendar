FROM golang:1.22.3-alpine3.19

RUN apk update && apk add bash

WORKDIR /app

COPY . .
COPY ./go.mod .
COPY ./go.sum .
COPY ./.air.toml .

RUN go mod download && \
    go build -o myapp ./cmd/main && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/air-verse/air@latest

# airを起動
CMD ["air", "-c", ".air.toml"]
