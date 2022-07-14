FROM golang:1.18-alpine as build-env
#FROM alpine:3.16 as build-env
#RUN apk add go>1.18 && go version

ARG APP_NAME=posts-go-api
ENV APP_NAME=$APP_NAME

WORKDIR /app
COPY . ./
#COPY .profile ./.profile

#RUN cd router && go mod tidy
#RUN go mod download

#RUN CGO_ENABLED=0 go build -v -o /$APP_NAME
RUN CGO_ENABLED=0 go build -o /posts-go-api

FROM alpine:3.14

ARG DEFAULT_PORT=8081
ARG APP_NAME=posts-go-api

ENV APP_NAME $APP_NAME
ENV PORT $DEFAULT_PORT

COPY .profile ./.profile
COPY --from=build-env /$APP_NAME .
#COPY --from=build-env /posts-go-api .

EXPOSE $DEFAULT_PORT

CMD ["./posts-go-api"]