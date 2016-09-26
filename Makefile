all: script

.PHONY: install
install:
	go get ./...

.PHONY: script
script: install
	go vet ./...
	go test ./...
