.PHONY: githooks install test goimports govet golint docker-build docker-push cover

githooks:
	ln -sf ../../githooks/pre-commit .git/hooks/pre-commit

install:
	go build -v ./...

test:
	go test ./...

cluster-test:
	go test --tags=cluster

goimports:
	go install golang.org/x/tools/cmd/goimports@latest
	goimports -w .
	go mod tidy

govet: goimports
	go vet ./...

golint: govet
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0
	golangci-lint run

API_DEFINITIONS_SHA=$(shell git log --oneline | grep Regenerated | head -n1 | cut -d ' ' -f 5)
docker-build:
	docker build -t twilio/twilio-go .
	docker tag twilio/twilio-go twilio/twilio-go:${GITHUB_TAG}
	docker tag twilio/twilio-go twilio/twilio-go:apidefs-${API_DEFINITIONS_SHA}
	docker tag twilio/twilio-go twilio/twilio-go:latest

docker-push:
	docker push twilio/twilio-go:${GITHUB_TAG}
	docker push twilio/twilio-go:apidefs-${API_DEFINITIONS_SHA}
	docker push twilio/twilio-go:latest

GO_DIRS = $(shell go list ./... | grep -v /rest/ | grep -v /form )
cover:
	go test ${GO_DIRS} -coverprofile coverage.out
	go test ${GO_DIRS} -json > test-report.out
