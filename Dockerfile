# syntax=docker/dockerfile:1

FROM golang:1.17-alpine
WORKDIR /playlist-server
COPY ./ /playlist-server
RUN go mod download
RUN go build -o ./playlist-server cmd/main.go
EXPOSE 8083
EXPOSE 8081
CMD ./playlist-server -p 8083

