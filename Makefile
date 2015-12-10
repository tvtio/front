build:
	go build

install:
	go install

test:
	go test -v ./...

cover:
	go test -v ./... -cover

lint:
	go vet ./...
	golint ./...

clean:
	go clean

deps:
	go get -u github.com/mgutz/ansi/cmd/ansi-mgutz
	go get -u github.com/julienschmidt/httprouter
	go get -u github.com/codegangsta/negroni
	go get -u github.com/phyber/negroni-gzip/gzip
	go get -u github.com/goincremental/negroni-sessions
	go get -u golang.org/x/oauth2
	go get -u github.com/go-sql-driver/mysql

dev-deps:
	go get -u github.com/stretchr/testify

docker:
	docker build -t tvtio -f docker/Dockerfile .

docker-run:
	docker run -p 80:8080 -ti --rm --name front -v `pwd`:/go/src/github.com/tvtio/front tvtio


.PHONY: test lint clean install docker docker-run
