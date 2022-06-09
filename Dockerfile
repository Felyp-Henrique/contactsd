FROM golang:1.18.3

WORKDIR /usr/local/go/src/contactsd

COPY * /usr/local/go/src/contactsd

RUN go get

ENTRYPOINT "go run /usr/local/go/src/contactsd/cmd/contactsd.go"
