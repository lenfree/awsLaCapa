DEFAULT_NAME := awslacapa
CONTAINER_PORT ?= 8080
DOCKER_IMAGE_NAME ?= $(DEFAULT_NAME)

project_name = awsLaCapa
package = github.com/lenfree/$(project_name)

all: release

.PHONY: install
install:
	go get -u github.com/smartystreets/goconvey
	go get github.com/beego/bee
	go get -v

.PHONY: script
script: install
	go vet ./...
	go test -v -race ./...
	go build -v -tags test ./...

.PHONY: test
test:
	go test -v -race ./...

.PHONY: build
build: image
	docker build -t $(DOCKER_IMAGE_NAME) .

.PHONY: run
run:
ifndef HOST_PORT
	$(error HOST_PORT is not set)
endif
	docker run -d -p $(HOST_PORT):$(CONTAINER_PORT) -v $(PWD):/root/.aws/ $(DOCKER_IMAGE_NAME)

.PHONY: image
image:
DOCKER_IMAGE_NAME ?= $(DEFAULT_NAME)

.PHONY: appstart
appstart:
	bee run

.PHONY: release
release:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/$(project_name)-linux-amd64 $(package)
	GOOS=darwin GOARCH=amd64 go build -o release/$(project_name)-darwin-amd64 $(package)

.PHONY: help
help:
	@echo "Usage: make <command>"
	@echo "  install       Install all required Go packages"
	@echo "  script        Install all rquired Go packages and run test"
	@echo "  test          Run test"
	@echo "  build         Build Docker container with name DOCKER_IMAGE_NAME"
	@echo "                  Default container name is awslacapa"
	@echo "  run           Run Docker container DOCKER_IMAGE_NAME with port"
	@echo "                  Default container is awslacapa"
	@echo "  appstart      Run the app and start web server"
