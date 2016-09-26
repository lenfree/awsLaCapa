all: script

.PHONY: install
install:
	go get -u github.com/smartystreets/goconvey
	go get -u -v $(go list -f '{{join .Imports "\n"}}' ./... | sort | uniq | grep -v golang-samples)

.PHONY: script
script: install
	go vet ./...
	go test -race ./...
