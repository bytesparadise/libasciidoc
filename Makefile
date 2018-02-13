# Makefile for the `libasciidoc` project

# tools
GIT_BIN_NAME := git
GIT_BIN := $(shell command -v $(GIT_BIN_NAME) 2> /dev/null)
DEP_BIN_NAME := dep
DEP_BIN := $(shell command -v $(DEP_BIN_NAME) 2> /dev/null)
GO_BIN_NAME := go
GO_BIN := $(shell command -v $(GO_BIN_NAME) 2> /dev/null)
EXTRA_PATH=$(shell dirname $(GO_BINDATA_BIN))

PROJECT_NAME=libasciidoc
PACKAGE_NAME := github.com/bytesparadise/libasciidoc
CUR_DIR=$(shell pwd)
TMP_PATH=$(CUR_DIR)/tmp
INSTALL_PREFIX=$(CUR_DIR)/bin
BINARY_PATH=$(INSTALL_PREFIX)/libasciidoc
VENDOR_DIR=vendor
SOURCE_DIR ?= .
SOURCES := $(shell find $(SOURCE_DIR) -path $(SOURCE_DIR)/vendor -prune -o -name '*.go' -print)
COVERPKGS := $(shell go list ./... | grep -v vendor | paste -sd "," -)
COMMIT=$(shell git rev-parse HEAD)
# GITUNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
GITUNTRACKEDCHANGES := $(shell git status --porcelain)
ifneq ($(GITUNTRACKEDCHANGES),)
COMMIT := $(COMMIT)-dirty
endif
BUILD_TIME=`date -u '+%Y-%m-%dT%H:%M:%SZ'`



# Pass in build time variables to main
LDFLAGS=-ldflags "-X ${PACKAGE_NAME}/about.Commit=${COMMIT} -X ${PACKAGE_NAME}/about.BuildTime=${BUILD_TIME}"

$(INSTALL_PREFIX):
# Build artifacts dir
	@mkdir -p $(INSTALL_PREFIX)

$(TMP_PATH):
	@mkdir -p $(TMP_PATH)


# Call this function with $(call log-info,"Your message")
define log-info =
@echo "INFO: $(1)"
endef

# If nothing was specified, run all targets as if in a fresh clone
.PHONY: all
## Default target - fetch dependencies, generate code and build.
all: prebuild-check get-deps generate build

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

.PHONY: get-deps 
## Download build dependencies.
get-deps: $(VENDOR_DIR) 

$(VENDOR_DIR): 
	$(DEP_BIN) ensure 

.PHONY: prebuild-checks
prebuild-checks: $(TMP_PATH) $(INSTALL_PREFIX) 
# Check that all tools where found
ifndef GIT_BIN
	$(error The "$(GIT_BIN_NAME)" executable could not be found in your PATH)
endif
ifndef DEP_BIN
	$(error The "$(DEP_BIN_NAME)" executable could not be found in your PATH)
endif

ifndef GO_BIN
	$(error The "$(GO_BIN_NAME)" executable could not be found in your PATH)
endif

.PHONY: generate
## generates the .go file based on the asciidoc grammar
generate:
	@pigeon ./parser/asciidoc-grammar.peg > ./parser/asciidoc_parser.go


.PHONY: get-deps test
## run all tests except in the 'vendor' package 
test: 
	@echo $(COVERPKGS)
	@ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --trace --race --compilers=2  --cover -coverpkg $(COVERPKGS)
