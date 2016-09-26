all: script

.PHONY: install
install:
	- go get ./...

.PHONY: script
script:
	- go vet ./...
	- go test ./...
