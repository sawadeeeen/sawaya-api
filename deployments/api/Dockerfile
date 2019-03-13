FROM golang:latest

ENV GOPATH $GOPATH:$HOME/go
EXPOSE 8080

RUN go get github.com/gin-gonic/gin
RUN go get github.com/sawadeeeen/sawaya-api

WORKDIR /go/src/github.com/sawadeeeen/sawaya-api

CMD ["go", "run", "server.go"]
