FROM golang:1.15.0-alpine AS build
RUN apk add --update --no-cache git
WORKDIR /go/src/github.com/traPtitech/NeoShowcase
COPY ./go.* ./
RUN go mod download
COPY . .

ARG APP_VERSION=dev
ARG APP_REVISION=local
RUN CGO_ENABLED=0 go build -o /neoshowcase -ldflags "-s -w -X main.version=$APP_VERSION -X main.revision=$APP_REVISION"

FROM alpine:3.12.0
WORKDIR /app

RUN apk add --update ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

EXPOSE 8080

COPY --from=build /neoshowcase ./

ENTRYPOINT ./neoshowcase
