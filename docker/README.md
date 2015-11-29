# tvt.io on Docker

## How to build

```
$ docker build -t tvtio -f docker/Dockerfile .
```

## How to run

```
$ docker run -ti --rm tvtio .
```

Exports containerized service from port 8080 to port 80 on the host machine.

Try: http://<docker-host>/
