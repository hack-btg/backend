FROM golang:latest AS build
WORKDIR /src
ENV GO111MODULE=on
COPY . .
RUN go mod download
RUN go build -o api ./application.go
EXPOSE 5001
ENTRYPOINT ./api