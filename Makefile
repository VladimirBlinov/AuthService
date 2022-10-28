.PHONY: build
build:
		go build -v ./cmd/server

.PHONY: test
test:
		go test -v -race -timeout 30s ./internal/...

.DEFAULT_GOAL := build