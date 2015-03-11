build: clean
	go build

test: clean build
	go test -v

deps: dist-clean
	go get -u github.com/moviewio/tmdb

clean:
	go clean

dist-clean:
	rm -rf pkg src

.PHONY: build test deps clean
