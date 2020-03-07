FROM golang:1.14-alpine

RUN mkdir -p /opt/benkyo/cmd && \
  mkdir -p $GOPATH/src/github.com/tanookiben/benkyo

WORKDIR $GOPATH/src/github.com/tanookiben/benkyo

COPY . .

RUN go build -o /opt/benkyo/cmd/api cmd/api/main.go

EXPOSE 8000

CMD ["/opt/benkyo/cmd/api"]
