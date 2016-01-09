build: tags
	go build

build-linux:
	GOOS=linux GOARCH=amd64 go build

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
	go get github.com/mgutz/ansi/cmd/ansi-mgutz
	go get github.com/julienschmidt/httprouter
	go get github.com/codegangsta/negroni
	go get github.com/phyber/negroni-gzip/gzip
	go get github.com/goincremental/negroni-sessions
	go get golang.org/x/oauth2
	go get github.com/repejota/kvson

dev-deps:
	go get github.com/golang/lint/golint
	go get github.com/stretchr/testify
	go get github.com/jstemmer/gotags

tags:
	gotags -tag-relative=true -R=true -sort=true -f="tags" -fields=+l .

docker:
	docker build -t docker_front .

docker-run: docker
	docker run -p 80:8080 -ti --rm --name front -v `pwd`:/go/src/github.com/tvtio/front docker_front
