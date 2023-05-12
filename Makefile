default: install

.PHONY: install
install: build
	sudo mv esc /usr/local/bin

.PHONY: build
build:
	go build -o esc .

.PHONY: test
test:
	go test -v ./...

.PHONY: format
format:
	go fmt ./...