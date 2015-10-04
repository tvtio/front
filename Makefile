build: clean
	go build

install: build
	go install

test: build
	go test -v

lint: build
	golint

clean:
	go clean

deps:
	go get -u github.com/julienschmidt/httprouter
	go get -u github.com/codegangsta/negroni
	go get -u github.com/phyber/negroni-gzip/gzip
	go get -u github.com/goincremental/negroni-sessions
	# Dev dependencies
	go get -u github.com/stretchr/testify

dist-clean: clean
	rm -rf pkg src bin

.PHONY: build test lint deps clean dist-clean
