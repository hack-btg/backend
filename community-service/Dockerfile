FROM golang:latest AS build
WORKDIR /src
ENV GO111MODULE=on
COPY . .
RUN go mod download
RUN go build -o api ./main.go
EXPOSE 5003
ENTRYPOINT ./api