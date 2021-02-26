.PHONY: install

install:
	go build -v ./...

test:
	go test ./...

goimports:
	go get golang.org/x/tools/cmd/goimports
	goimports -w twilio
