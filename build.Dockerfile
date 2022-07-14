# Build Stage

#FROM golang:1.18-alpine as build-env
FROM alpine:3.16 as build-env
RUN apk add go>1.18 # 1.18.2-r0
RUN go version

ARG APP_NAME=posts-go-api
ENV APP_NAME=$APP_NAME

WORKDIR /app
COPY . ./
#COPY .profile ./.profile

#RUN go version
#
#RUN cd router && go mod tidy
#
#RUN go mod download
##RUN CGO_ENABLED=0 go build -v -o /$APP_NAME
#RUN CGO_ENABLED=0 go build -o /posts-go-api