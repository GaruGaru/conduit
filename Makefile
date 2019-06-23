VERSION=$(shell git rev-parse --short HEAD)
DOCKER_IMAGE=garugaru/conduit

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
	docker build -t ${DOCKER_IMAGE}:${VERSION} -t ${DOCKER_IMAGE}:latest .

docker-push: docker-build
	docker push ${DOCKER_IMAGE}:${VERSION}
	docker push ${DOCKER_IMAGE}:latest