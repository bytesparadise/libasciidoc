# By default the project should be build under GOPATH/src/github.com/<orgname>/<reponame>
GO_PACKAGE_ORG_NAME ?= $(shell basename $$(dirname $$PWD))
GO_PACKAGE_REPO_NAME ?= $(shell basename $$PWD)
GO_PACKAGE_PATH ?= github.com/${GO_PACKAGE_ORG_NAME}/${GO_PACKAGE_REPO_NAME}


ifeq ($(OS),Windows_NT)
BINARY_PATH=$(INSTALL_PREFIX)/libasciidoc.exe
else
BINARY_PATH=$(INSTALL_PREFIX)/libasciidoc
endif

# Call this function with $(call log-info,"Your message")
define log-info =
@echo "INFO: $(1)"
endef

CUR_DIR=$(shell pwd)
INSTALL_PREFIX=$(CUR_DIR)/bin

$(INSTALL_PREFIX):
# Build artifacts dir
	@mkdir -p $(INSTALL_PREFIX)

.PHONY: prebuild-checks
## Check that all tools where found
prebuild-checks: $(INSTALL_PREFIX)

.PHONY: install-pigeon
## Install development tools.
install-pigeon:
	@go install -v github.com/mna/pigeon

.PHONY: generate
## generate the .go file based on the asciidoc grammar
generate: install-pigeon
	@if [ "pkg/parser/parser.go" -ot "pkg/parser/parser.peg" ]; then \
		echo "generating the parser..."; \
		pigeon ./pkg/parser/parser.peg > ./pkg/parser/parser.go; \
		echo "done"; \
	else \
		echo "no need to regenerate the parser."; \
	fi;

.PHONY: verify-parser
## verify that the parser was built with the latest version of pigeon, using the `optimize-grammar` option
verify-parser: prebuild-checks clean
ifneq ($(shell git diff --quiet pkg/parser/parser.go; echo $$?), 0)
	@git diff pkg/parser/parser.go
	$(error "pkg/parser/parser.go is uncommited or was generated with an older version of 'mna/pigeon' or without the '-optimize-parser' option enabled.")
else
	@echo "generated parser is ok"
endif

.PHONY: generate-optimized
## generate the .go file based on the asciidoc grammar
generate-optimized: install-pigeon
	@if [ "pkg/parser/parser.go" -ot "pkg/parser/parser.peg" ]; then \
		echo "generating the parser (optimized)..."; \
		go generate ./...; \
		echo "done"; \
	else \
		echo "no need to regenerate the parser."; \
	fi;


.PHONY: build
## build the binary executable from CLI
build: prebuild-checks verify-parser generate-optimized
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

PARSER_DIFF_STATUS :=

.PHONY: install
## installs the binary executable in the $GOPATH/bin directory
install: build
	@cp $(BINARY_PATH) $(GOPATH)/bin

