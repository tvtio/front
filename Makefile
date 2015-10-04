build:
	go build

install:
	go install

test:
	go test -v ./...

cover:
	go test -v ./... -cover

lint:
	golint

clean:
	go clean

deps:
	go get -u github.com/julienschmidt/httprouter
	go get -u github.com/codegangsta/negroni
	go get -u github.com/phyber/negroni-gzip/gzip
	go get -u github.com/goincremental/negroni-sessions
	go get -u golang.org/x/oauth2
	# Dev dependencies
	go get -u github.com/stretchr/testify

dist-clean: clean
	rm -rf pkg src bin

.PHONY: build test lint deps clean dist-clean
