.PHONY: install test goimports govet golint githooks

install:
	go build -v ./...

test:
	go test -v ./...

goimports:
	go get golang.org/x/tools/cmd/goimports
	goimports -w .

govet: goimports
	go vet ./...

golint: govet
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0
	golangci-lint run --new-from-rev=main --enable-all --exclude-use-default=false -D errcheck -D gomnd

githooks:
	cp githooks/pre-commit `git rev-parse --git-dir`/hooks/pre-commit
	chmod +x `git rev-parse --git-dir`/hooks/pre-commit
