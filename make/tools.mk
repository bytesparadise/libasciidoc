# tools
VENDOR_DIR=vendor
SOURCE_DIR ?= .
COVERPKGS := $(shell go list ./... | grep -v vendor | paste -sd "," -)
DEVTOOLS=\
	github.com/mna/pigeon \
	github.com/onsi/ginkgo/ginkgo \
	github.com/sozorogami/gover \
	github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: install-devtools
## Install development tools.
install-devtools:
	@go mod download
	@go install -v $(DEVTOOLS)
