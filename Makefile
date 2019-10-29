# Makefile for the `libasciidoc` project

# tools
CUR_DIR=$(shell pwd)
INSTALL_PREFIX=$(CUR_DIR)/bin
VENDOR_DIR=vendor
SOURCE_DIR ?= .
COVERPKGS := $(shell go list ./... | grep -v vendor | paste -sd "," -)

DEVTOOLS=\
				github.com/mna/pigeon \
				github.com/onsi/ginkgo/ginkgo \
				github.com/sozorogami/gover

ifeq ($(OS),Windows_NT)
BINARY_PATH=$(INSTALL_PREFIX)/libasciidoc.exe
else
BINARY_PATH=$(INSTALL_PREFIX)/libasciidoc
endif

# Call this function with $(call log-info,"Your message")
define log-info =
@echo "INFO: $(1)"
endef


.PHONY: help
# Based on https://gist.github.com/rcmachado/af3db315e31383502660
## Display this help text.
help:/
	$(info Available targets)
	$(info -----------------)
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		helpCommand = substr($$1, 0, index($$1, ":")-1); \
		if (helpMessage) { \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			gsub(/##/, "\n                                     ", helpMessage); \
		} else { \
			helpMessage = "(No documentation)"; \
		} \
		printf "%-35s - %s\n", helpCommand, helpMessage; \
		lastLine = "" \
	} \
	{ hasComment = match(lastLine, /^## (.*)/); \
          if(hasComment) { \
            lastLine=lastLine$$0; \
	  } \
          else { \
	    lastLine = $$0 \
          } \
        }' $(MAKEFILE_LIST)

.PHONY: install-devtools
## Install development tools.
install-devtools:
	@go mod download
	@go install -v $(DEVTOOLS)

$(INSTALL_PREFIX):
# Build artifacts dir
	@mkdir -p $(INSTALL_PREFIX)

.PHONY: prebuild-checks
## Check that all tools where found
prebuild-checks: $(INSTALL_PREFIX)

.PHONY: generate
## generate the .go file based on the asciidoc grammar
generate: prebuild-checks
	@echo "generating the parser..."
	@pigeon ./pkg/parser/parser.peg > ./pkg/parser/parser.go

.PHONY: generate-optimized
## generate the .go file based on the asciidoc grammar
generate-optimized:
	@echo "generating the parser (optimized)..."
	@pigeon -optimize-parser \
		-alternate-entrypoints PreflightAsciidocDocument,PreflightAsciidocDocumentWithinDelimitedBlock,PreflightTextDocument,DocumentBlock,InlineElementsWithoutSubtitution,FileLocation,IncludedFileLine \
		-o ./pkg/parser/parser.go ./pkg/parser/parser.peg

.PHONY: test
## run all tests excluding fixtures and vendored packages
test: generate-optimized
	@echo $(COVERPKGS)
	@ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --trace --race --compilers=0

.PHONY: test-with-coverage
## run all tests excluding fixtures and vendored packages
test-with-coverage: generate-optimized
	@echo $(COVERPKGS)
	@ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --trace --race --compilers=0  --cover -coverpkg $(COVERPKGS)

.PHONY: test-fixtures
## run all fixtures tests
test-fixtures: generate-optimized
	@ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --trace --race --compilers=2 -tags=fixtures --focus=fixtures

.PHONY: bench-parser
## run the benchmarks on the parser
bench-parser: generate-optimized
	$(eval GIT_BRANCH:=$(shell git rev-parse --abbrev-ref HEAD))
	go test -run="XXX" -bench=. -benchmem -count=10 \
		github.com/bytesparadise/libasciidoc/pkg/parser | \
		tee ./tmp/bench-$(GIT_BRANCH).txt

.PHONY: build
## build the binary executable from CLI
build: $(INSTALL_PREFIX) generate-optimized
	$(eval BUILD_COMMIT:=$(shell git rev-parse --short HEAD))
	$(eval BUILD_TAG:=$(shell git tag --contains $(BUILD_COMMIT)))
	$(eval BUILD_TIME:=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ'))
	@echo "building $(BINARY_PATH) (commit:$(BUILD_COMMIT) / tag:$(BUILD_TAG) / time:$(BUILD_TIME))"
	@go build -ldflags \
	  " -X github.com/bytesparadise/libasciidoc.BuildCommit=$(BUILD_COMMIT)\
	    -X github.com/bytesparadise/libasciidoc.BuildTag=$(BUILD_TAG) \
	    -X github.com/bytesparadise/libasciidoc.BuildTime=$(BUILD_TIME)" \
	  -o $(BINARY_PATH) \
	  cmd/libasciidoc/*.go

.PHONY: lint
## run golangci-lint against project
lint:
	@golangci-lint run -E gofmt,golint,megacheck,misspell ./...

PARSER_DIFF_STATUS :=

.PHONY: verify-parser
## verify that the parser was built with the latest version of pigeon, using the `optimize-grammar` option
verify-parser: prebuild-checks
ifneq ($(shell git diff --quiet pkg/parser/parser.go; echo $$?), 0)
	$(error "parser was generated with an older version of 'mna/pigeon' or without the '-optimize' option(s).")
else
	@echo "parser is ok"
endif

.PHONY: install
## installs the binary executable in the $GOPATH/bin directory
install: install-devtools build
	@cp $(BINARY_PATH) $(GOPATH)/bin

.PHONY: quick-install
## installs the binary executable in the $GOPATH/bin directory without prior tools setup and parser generation
quick-install:
	$(eval BUILD_COMMIT:=$(shell git rev-parse --short HEAD))
	$(eval BUILD_TAG:=$(shell git tag --contains $(BUILD_COMMIT)))
	$(eval BUILD_TIME:=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ'))
	@echo "building $(BINARY_PATH) (commit:$(BUILD_COMMIT) / tag:$(BUILD_TAG) / time:$(BUILD_TIME))"
	@go build -ldflags \
	  " -X github.com/bytesparadise/libasciidoc.BuildCommit=$(BUILD_COMMIT)\
	    -X github.com/bytesparadise/libasciidoc.BuildTag=$(BUILD_TAG) \
	    -X github.com/bytesparadise/libasciidoc.BuildTime=$(BUILD_TIME)" \
	  -o $(BINARY_PATH) \
	  cmd/libasciidoc/*.go
	@cp $(BINARY_PATH) $(GOPATH)/bin
