VERSION=$(shell git rev-parse --short HEAD)
DOCKER_IMAGE=garugaru/conduit:${VERSION}

fmt:
	go fmt ./...

deps:
	go mod vendor
	go mod verify

test:
	go test ./...

build: fmt deps
	go build -o ${BIN_OUTPUT}

docker-build:
	docker build -t ${DOCKER_IMAGE} .

docker-push: docker-build
	docker push ${DOCKER_IMAGE}