# Makefile for the `libasciidoc` project

# tools
GIT_BIN_NAME := git
GIT_BIN := $(shell command -v $(GIT_BIN_NAME) 2> /dev/null)
GLIDE_BIN_NAME := glide
GLIDE_BIN := $(shell command -v $(GLIDE_BIN_NAME) 2> /dev/null)
GO_BIN_NAME := go
GO_BIN := $(shell command -v $(GO_BIN_NAME) 2> /dev/null)
EXTRA_PATH=$(shell dirname $(GO_BINDATA_BIN))
GOCOV_BIN=$(VENDOR_DIR)/github.com/axw/gocov/gocov/gocov
GOCOVMERGE_BIN=$(VENDOR_DIR)/github.com/wadey/gocovmerge/gocovmerge
GOLINT_DIR=$(VENDOR_DIR)/github.com/golang/lint/golint
GOLINT_BIN=$(GOLINT_DIR)/golint
GOCYCLO_DIR=$(VENDOR_DIR)/github.com/fzipp/gocyclo
GOCYCLO_BIN=$(GOCYCLO_DIR)/gocyclo

PROJECT_NAME=libasciidoc
PACKAGE_NAME := github.com/bytesparadise/libasciidoc
CUR_DIR=$(shell pwd)
TMP_PATH=$(CUR_DIR)/tmp
INSTALL_PREFIX=$(CUR_DIR)/bin
BINARY_PATH=$(INSTALL_PREFIX)/libasciidoc
VENDOR_DIR=vendor
SOURCE_DIR ?= .
SOURCES := $(shell find $(SOURCE_DIR) -path $(SOURCE_DIR)/vendor -prune -o -name '*.go' -print)
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
all: prebuild-check install-deps generate build

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

.PHONY: check-go-format
## Exists with an error if there are files whose formatting differs from gofmt's
check-go-format: prebuild-check
	@gofmt -s -l ${SOURCES} 2>&1 \
		| tee /tmp/gofmt-errors \
		| read \
	&& echo "ERROR: These files differ from gofmt's style (run 'make format-go-code' to fix this):" \
	&& cat /tmp/gofmt-errors \
	&& exit 1 \
	|| true

.PHONY: analyze-go-code
## Run a complete static code analysis using the following tools: golint, gocyclo and go-vet.
analyze-go-code: golint gocyclo govet

## Run gocyclo analysis over the code.
golint: $(GOLINT_BIN)
	$(info >>--- RESULTS: GOLINT CODE ANALYSIS ---<<)
	@$(foreach d,$(GOANALYSIS_DIRS),$(GOLINT_BIN) $d 2>&1 | grep -vEf .golint_exclude;)

$(GOLINT_BIN):
	cd $(VENDOR_DIR)/github.com/golang/lint/golint && go build -v

# Build go tool to analysis the code

## Run gocyclo analysis over the code.
gocyclo: $(GOCYCLO_BIN)
	$(info >>--- RESULTS: GOCYCLO CODE ANALYSIS ---<<)
	@$(foreach d,$(GOANALYSIS_DIRS),$(GOCYCLO_BIN) -over 15 $d | grep -vEf .golint_exclude;)

$(GOCYCLO_BIN):
	cd $(VENDOR_DIR)/github.com/fzipp/gocyclo && go build -v

## Run go vet analysis over the code.
govet:
	$(info >>--- RESULTS: GO VET CODE ANALYSIS ---<<)
	@$(foreach d,$(GOANALYSIS_DIRS),go tool vet --all $d/*.go 2>&1;)

.PHONY: format-go-code
## Formats any go file that differs from gofmt's style
format-go-code: prebuild-check
	@gofmt -s -l -w ${SOURCES}

.PHONY: deps 
## Download build dependencies.
deps: $(VENDOR_DIR) 

$(VENDOR_DIR): glide.lock glide.yaml
	$(GLIDE_BIN) update 
	touch $(VENDOR_DIR)

.PHONY: prebuild-check
prebuild-check: $(TMP_PATH) $(INSTALL_PREFIX) 
# Check that all tools where found
ifndef GIT_BIN
	$(error The "$(GIT_BIN_NAME)" executable could not be found in your PATH)
endif
ifndef GLIDE_BIN
	$(error The "$(GLIDE_BIN_NAME)" executable could not be found in your PATH)
endif

ifndef GO_BIN
	$(error The "$(GO_BIN_NAME)" executable could not be found in your PATH)
endif

# Keep this "clean" target here at the bottom
.PHONY: clean
## Runs all clean-* targets.
clean: $(CLEAN_TARGETS)

CLEAN_TARGETS += clean-artifacts
.PHONY: clean-artifacts
## Removes the ./bin directory.
clean-artifacts:
	@rm -rf $(INSTALL_PREFIX)

CLEAN_TARGETS += clean-object-files
.PHONY: clean-object-files
## Runs go clean to remove any executables or other object files.
clean-object-files:
	@go clean ./...

CLEAN_TARGETS += clean-vendor
.PHONY: clean-vendor
## Removes the ./vendor directory.
clean-vendor:
	@rm -rf $(VENDOR_DIR)

CLEAN_TARGETS += clean-glide-cache
.PHONY: clean-glide-cache
## Removes the ./glide directory.
clean-glide-cache:
	@rm -rf ./.glide

.PHONY: generate
## generates the .go file based on the asciidoc grammar
generate:
	@pigeon ./parser/asciidoc-grammar.peg > ./parser/asciidoc_parser.go


.PHONY: build
## builds the application.
build: prebuild-check clean-artifacts $(BINARY_PATH)

# $(BINARY_PATH): $(SOURCES)
$(BINARY_PATH): 
	@rm -f ${BINARY_PATH}
	@echo Building from $(COMMIT)
	@go build -v ${LDFLAGS} -o ${BINARY_PATH}
    