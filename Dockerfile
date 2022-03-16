# syntax=docker/dockerfile:1

FROM golang:onbuild
WORKDIR /playlist-server
COPY . .
RUN go mod download
RUN go build -o ./playlist-server cmd/main.go
EXPOSE 8083
EXPOSE 8081
CMD ./playlist-server

