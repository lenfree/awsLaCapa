DEFAULT_NAME := awslacapa
CONTAINER_PORT ?= 8080
DOCKER_IMAGE_NAME ?= $(DEFAULT_NAME)

project_name = awsLaCapa
package = github.com/lenfree/$(project_name)

all: script

.PHONY: install
install:
	go get -u github.com/smartystreets/goconvey
	go get github.com/beego/bee
	go get -v

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
	docker run -d -p $(HOST_PORT):$(CONTAINER_PORT) -v $(PWD):/root/.aws/ $(DOCKER_IMAGE_NAME)

.PHONY: image
image:
DOCKER_IMAGE_NAME ?= $(DEFAULT_NAME)

.PHONY: appstart
appstart: aws-credentials
	bee run

# This is dangerous. One might accidentally push to public repo. I'd prefer
# using environment varibles instead
.PHONY: aws-credentials
aws-credentials:
	cp credentials /root/.aws

.PHONY: release
release:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/$(DEFAULT_NAME)-linux-amd64 $(package)
	GOOS=darwin GOARCH=amd64 go build -o release/$(DEFAULT_NAME)-darwin-amd64 $(package)

.PHONY: help
help:
	@echo "Usage: make <command>"
	@echo "  install       Install all required Go packages"
	@echo "  script        Install all rquired Go packages and run test"
	@echo "  build         Build Docker container with name DOCKER_IMAGE_NAME"
	@echo "                  Default container name is awslacapa"
	@echo "  run           Run Docker container DOCKER_IMAGE_NAME with port"
	@echo "                  Default container is awslacapa"
	@echo "  appstart      Run the app and start web server"
