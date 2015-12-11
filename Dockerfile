FROM golang:latest
MAINTAINER Raül Pérez <repejota@gmail.com>

ENV GOPATH /go
ENV GOBIN /go/bin
ENV PATH /usr/local/go/bin:/go/bin:$PATH

ADD . $GOPATH/src/github.com/tvtio/front
WORKDIR $GOPATH/src/github.com/tvtio/front
RUN make deps
RUN make install

VOLUME ["/cache"]

EXPOSE 80:8080

CMD ["bash"]
