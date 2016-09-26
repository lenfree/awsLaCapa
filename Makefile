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
build:
	docker buid -t awsLaCapa .

.PHONY: run
run: docker-port
	docker run -d -p PORT:PORT -v $(PWD):/root/.aws/ awsLaCapa .

.PHONY: docker-port
docker-port:
ifndef PORT
	$(error docker PORT is not specified)
endif
