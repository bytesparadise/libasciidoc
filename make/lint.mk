.PHONY: install-golangci-lint
## Install development tools.
install-golangci-lint:
	@go install -v github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: lint
## run golangci-lint against project
lint: install-golangci-lint
	@golangci-lint run -E gofmt,golint,megacheck,misspell ./...
