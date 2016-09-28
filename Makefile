DEFAULT_NAME := awslacapa
CONTAINER_PORT ?= 8080

all: script

.PHONY: install
install:
	go get -u github.com/smartystreets/goconvey
	go get -u -v $(go list -f '{{join .Imports "\n"}}' ./... | sort | uniq | grep -v golang-samples)

.PHONY: script
script: install
	go vet ./...
	go test -v -race ./...

.PHONY: build
build: image
	docker build -t $(DOCKER_IMAGE_NAME) .

.PHONY: run
run:
ifndef HOST_PORT
	$(error HOST_PORT is not set)
endif
	docker run -d -p $(PORT):$(CONTAINER_PORT) -v $(PWD):/root/.aws/ awsLaCapa .

.PHONY: image
image:
DOCKER_IMAGE_NAME ?= $(DEFAULT_NAME)


.PHONY: help
help:
	@echo "Usage: make <command>"
	@echo "  install       Install all required Go packages"
	@echo "  script        Install all rquired Go packages and run test"
	@echo "  build         Build Docker container with name DOCKER_IMAGE_NAME"
	@echo "                  Default container name is awslacapa"
	@echo "  run           Run Docker container DOCKER_IMAGE_NAME with port"
	@echo "                  Default container is awslacapa"
