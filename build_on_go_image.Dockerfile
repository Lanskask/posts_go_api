FROM golang:1.18-alpine as build-env

ARG APP_NAME=posts-go-api
ARG DEFAULT_PORT=8081
ENV APP_NAME=$APP_NAME
ENV DEFAULT_PORT=$DEFAULT_PORT

WORKDIR /app
COPY go.mod ./go.mod
RUN go mod download

COPY . ./

EXPOSE $DEFAULT_PORT

RUN CGO_ENABLED=0 go build -v -o ./$APP_NAME

CMD ./$APP_NAME