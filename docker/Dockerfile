FROM golang:latest

ENV GOPATH /go
ENV GOBIN /go/bin
ENV PATH /usr/local/go/bin:/go/bin:$PATH

VOLUME $GOPATH/src/github.com/tvtio/front
WORKDIR $GOPATH/src/github.com/tvtio/front

RUN go get -u github.com/mgutz/ansi/cmd/ansi-mgutz
RUN go get -u github.com/julienschmidt/httprouter
RUN go get -u github.com/codegangsta/negroni
RUN go get -u github.com/phyber/negroni-gzip/gzip
RUN go get -u github.com/goincremental/negroni-sessions
RUN go get -u golang.org/x/oauth2
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/repejota/kvson
RUN go get -u github.com/stretchr/testify

ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
EXPOSE 80:8080
