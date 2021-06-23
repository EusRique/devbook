FROM golang:alpine

WORKDIR /devbook

ADD . .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -command="./devbook"