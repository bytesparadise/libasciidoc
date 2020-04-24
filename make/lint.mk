.PHONY: lint
## run golangci-lint against project
lint:
	@golangci-lint run -E gofmt,golint,megacheck,misspell ./...
