.DEFAULT_GOAL := docker-run

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

docker:
	docker build -t tvtio -f docker/Dockerfile .

docker-run:
	docker run -p 80:8080 -ti --rm --name front -v `pwd`:/go/src/github.com/tvtio/front tvtio


.PHONY: test lint clean install docker docker-run
