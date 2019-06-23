VERSION=$(shell git rev-parse --short HEAD)
DOCKER_IMAGE=garugaru/conduit:${VERSION}

docker-build:
	docker build -t ${DOCKER_IMAGE} .
