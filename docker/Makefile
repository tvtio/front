build:
	docker build -t tvtio .

run:
	docker run -p 80:8080 -ti --rm --name front -v `pwd`/..:/go/src/github.com/tvtio/front tvtio
