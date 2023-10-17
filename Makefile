default: test

.PHONY: build
build:
	go build -o esc .

.PHONY: test
test:
	go test -v ./...

.PHONY: format
format:
	go fmt ./...
