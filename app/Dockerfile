FROM golang:1.15
LABEL maintainer "akeno"

WORKDIR /go/src
COPY ./go.mod .
ENV GO111MODULE=on
RUN go mod download
EXPOSE 8080
CMD ["go", "run", "/go/src/app/main.go"]

