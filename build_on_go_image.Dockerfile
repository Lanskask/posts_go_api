FROM golang:1.18-alpine as build-env
#FROM alpine:3.16 as build-env
#RUN apk add go>1.18 && go version # 1.18.2-r0

ARG APP_NAME=posts-go-api
ARG DEFAULT_PORT=8081
ENV APP_NAME=$APP_NAME
ENV DEFAULT_PORT=$DEFAULT_PORT

WORKDIR /app
COPY . ./

#RUN go mod download
#

EXPOSE $DEFAULT_PORT

RUN CGO_ENABLED=0 go build -v -o ./$APP_NAME
#RUN CGO_ENABLED=0 go build -o ./posts-go-api
CMD ./$APP_NAME