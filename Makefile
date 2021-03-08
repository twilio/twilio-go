.PHONY: install test goimports govet golint githooks

install:
	go build -v ./...

test:
	go test ./...

goimports:
	go get golang.org/x/tools/cmd/goimports
	goimports -w twilio

govet: goimports
	go vet

golint: govet
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.37.0
	golangci-lint run

githooks:
	cp githooks/pre-commit `git rev-parse --git-dir`/hooks/pre-commit
	chmod +x `git rev-parse --git-dir`/hooks/pre-commit
