FROM golang:latest
MAINTAINER Raül Pérez <repejota@gmail.com>

ENV GOPATH /go
ENV GOBIN /go/bin
ENV PATH /usr/local/go/bin:/go/bin:$PATH

ADD front $GOBIN/front

EXPOSE 8080

CMD ["front"]
