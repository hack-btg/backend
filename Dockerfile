FROM golang:latest AS build
WORKDIR /src
COPY . .
RUN go build -o api ./cmd/api/main.go
EXPOSE 8888
ENTRYPOINT ./api