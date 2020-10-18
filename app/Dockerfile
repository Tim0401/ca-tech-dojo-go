FROM golang:1.15-alpine3.12 as build

LABEL maintainer="https://github.com/Tim0401"

WORKDIR /go/app

COPY . .

ENV GO111MODULE=off

RUN set -eux && \
  apk update && \
  apk add --no-cache git curl && \
  go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

ENV GO111MODULE on
# ENV PATH $PATH:/go/bin

RUN set -eux && \
  go build -o main ./main.go

# CMD ["./main"]

# FROM alpine:3.11

# WORKDIR /app

# COPY --from=build /go/app/main .

# RUN set -x && \
#   addgroup go && \
#   adduser -D -G go go && \
#   chown -R go:go /app/main

# CMD ["./main"]