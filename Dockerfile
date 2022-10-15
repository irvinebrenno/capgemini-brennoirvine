# Imagem do goserver
FROM golang:latest AS builder

LABEL maintainer="Brenno Irvine"

WORKDIR /go/src/capgemini-brennoirvine

ENV GOPATH=/app

COPY . /go/src/capgemini-brennoirvine

RUN go build main.go

ENTRYPOINT ["./main"]

EXPOSE 8080