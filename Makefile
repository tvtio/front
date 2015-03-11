build: clean
	go build

test: build
	go test -v

lint: build
	golint

clean:
	go clean

deps:
	# Go tools
	go get -u github.com/nsf/gocode
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/rogpeppe/godef
	go get -u golang.org/x/tools/cmd/oracle
	go get -u golang.org/x/tools/cmd/gorename
	go get -u github.com/golang/lint/golint
	go get -u github.com/kisielk/errcheck
	go get -u github.com/jstemmer/gotags
	# Dependencies
	go get -u github.com/moviewio/tmdb
	# Dev dependencies
	go get -u github.com/stretchr/testify

dist-clean: clean
	rm -rf pkg src bin

.PHONY: build test lint deps clean dist-clean
