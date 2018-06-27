# Makefile for the `libasciidoc` project

# tools
GIT_BIN_NAME := git
GIT_BIN := $(shell command -v $(GIT_BIN_NAME) 2> /dev/null)
DEP_BIN_NAME := dep
DEP_BIN := $(shell command -v $(DEP_BIN_NAME) 2> /dev/null)
PIGEON_BIN_NAME := pigeon
PIGEON_BIN := $(shell command -v $(PIGEON_BIN_NAME) 2> /dev/null)
GO_BIN_NAME := go
GO_BIN := $(shell command -v $(GO_BIN_NAME) 2> /dev/null)
EXTRA_PATH=$(shell dirname $(GO_BINDATA_BIN))

CUR_DIR=$(shell pwd)
TMP_PATH=$(CUR_DIR)/tmp
INSTALL_PREFIX=$(CUR_DIR)/bin
VENDOR_DIR=vendor
SOURCE_DIR ?= .
SOURCES := $(shell find $(SOURCE_DIR) -path $(SOURCE_DIR)/vendor -prune -o -name '*.go' -print)
COVERPKGS := $(shell go list ./... | grep -v vendor | paste -sd "," -)

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

.PHONY: deps 
## Download build dependencies.
deps: $(VENDOR_DIR) 

$(VENDOR_DIR): 
	$(DEP_BIN) ensure 

$(INSTALL_PREFIX):
# Build artifacts dir
	@mkdir -p $(INSTALL_PREFIX)

$(TMP_PATH):
	@mkdir -p $(TMP_PATH)

.PHONY: prebuild-checks
prebuild-checks: $(TMP_PATH) $(INSTALL_PREFIX) 
## Check that all tools where found
ifndef GIT_BIN
	$(error The "$(GIT_BIN_NAME)" executable could not be found in your PATH)
endif
ifndef DEP_BIN
	$(error The "$(DEP_BIN_NAME)" executable could not be found in your PATH)
endif
ifndef PIGEON_BIN
	$(error The "$(PIGEON_BIN_NAME)" executable could not be found in your PATH)
endif
ifndef GO_BIN
	$(error The "$(GO_BIN_NAME)" executable could not be found in your PATH)
endif

.PHONY: generate
## generate the .go file based on the asciidoc grammar
generate: prebuild-checks
	@echo "generating the parser..."
	@pigeon ./pkg/parser/asciidoc-grammar.peg > ./pkg/parser/asciidoc_parser.go

.PHONY: generate-optimized
## generate the .go file based on the asciidoc grammar
generate-optimized:
	@echo "generating the parser..."
	@pigeon -optimize-grammar ./pkg/parser/asciidoc-grammar.peg > ./pkg/parser/asciidoc_parser.go


.PHONY: test
## run all tests except in the 'vendor' package 
test: deps generate-optimized
	@echo $(COVERPKGS)
	@ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --trace --race --compilers=2  --cover -coverpkg $(COVERPKGS)

.PHONY: build
## build the binary executable from CLI
build: $(INSTALL_PREFIX) deps generate
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

.PHONY: verify-parser
## verify that the parser was built with the latest version of pigeon, using the `optimize-grammar` option
verify-parser: prebuild-checks
ifneq ($(shell git diff --quiet pkg/parser/asciidoc_parser.go; echo $$?), 0)
	$(error "parser was generated with an older version of 'mna/pigeon' or without the '-optimize' option(s).")
else 
	@echo "parser is ok"
endif
	

.PHONY: install
## installs the binary executable in the $GOPATH/bin directory
install: build
	@cp $(BINARY_PATH) $(GOPATH)/bin