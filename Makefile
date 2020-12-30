.PHONY: install

install:
	go build -v ./...

test:
	go test ./...
