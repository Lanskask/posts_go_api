FROM golang:1.18-alpine as build-env

ARG APP_NAME=posts-go-api
ENV APP_NAME=$APP_NAME

WORKDIR /app
COPY go.mod ./go.mod
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o /posts-go-api

FROM alpine:3.14

ARG DEFAULT_PORT=8081
ARG APP_NAME=posts-go-api

ENV APP_NAME $APP_NAME
ENV PORT $DEFAULT_PORT

WORKDIR /app
COPY .profile ./.profile
COPY --from=build-env /$APP_NAME .

EXPOSE $DEFAULT_PORT

CMD ["./posts-go-api"]