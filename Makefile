.PHONY: githooks install test goimports govet golint

githooks:
	ln -sf ../../githooks/pre-commit .git/hooks/pre-commit

install: githooks
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
	golangci-lint run
	go mod tidy
