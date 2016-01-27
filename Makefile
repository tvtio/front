build:
	go build

build-linux:
	GOOS=linux GOARCH=amd64 go build

install:
	go install

test:
	go test -v -race ./...

cover:
	go test -v -race  ./... -coverprofile=coverage.out
	goveralls -v -coverprofile=coverage.out -service=circle-ci -repotoken=8iNDEAQvVLpWexMuIMz9Nd8CXrVhLMJ5z
	rm coverage.out

lint:
	go vet ./...
	golint ./...

clean:
	go clean

deps: dev-deps
	go get -u github.com/mgutz/ansi/cmd/ansi-mgutz
	go get -u github.com/julienschmidt/httprouter
	go get -u github.com/codegangsta/negroni
	go get -u github.com/phyber/negroni-gzip/gzip
	go get -u github.com/goincremental/negroni-sessions
	go get -u golang.org/x/oauth2
	go get -u github.com/repejota/cache
	go get -u github.com/repejota/logger
	go get -u github.com/tvtio/tmdb

dev-deps:
	go get -u github.com/golang/lint/golint
	go get -u github.com/stretchr/testify
	go get -u github.com/jstemmer/gotags

tags:
	gotags -tag-relative=true -R=true -sort=true -f="tags" -fields=+l .

docker:
	docker build -t docker_front .

docker-run: docker
	docker run -p 80:8080 -ti --rm --name front -v `pwd`:/go/src/github.com/tvtio/front docker_front
