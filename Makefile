build: clean
	go build

test: build
	go test -v

lint: build
	golint

clean:
	go clean

deps:
	go get -u github.com/julienschmidt/httprouter
	# Dev dependencies
	go get -u github.com/stretchr/testify

dist-clean: clean
	rm -rf pkg src bin

.PHONY: build test lint deps clean dist-clean
