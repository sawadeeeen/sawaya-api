FROM golang:latest

ENV GOPATH $GOPATH:$HOME/go
EXPOSE 8080

RUN go get github.com/gin-gonic/gin
RUN go get github.com/sawadeeeen/sawaya-api

WORKDIR /usr/local/bin
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v3.2.0/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate.linux-amd64 migrate

WORKDIR /go/src/github.com/sawadeeeen/sawaya-api

CMD ["go", "run", "server.go"]
