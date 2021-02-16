.PHONY: install

install:
	go build -v ./...

test:
	go test ./...

fmt:
	gofmt -w twilio/*
