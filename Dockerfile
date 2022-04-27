#syntax=docker/dockerfile:1

FROM golang:1.18.1-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go clean
RUN go mod tidy
RUN go mod download

COPY *.go ./

RUN go build -o /auth-with-go-mysql

CMD ["/auth-with-go-mysql"]


