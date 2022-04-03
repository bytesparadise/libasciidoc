GIT_COMMIT_ID_SHORT := $(shell git rev-parse --short HEAD)
ifneq ($(shell git status --porcelain),)
       GIT_COMMIT_ID_SHORT := $(GIT_COMMIT_ID_SHORT)-dirty
endif

REF_NAME := $(shell git symbolic-ref -q --short HEAD || git describe --tags --exact-match)

BUILD_TIME = `date -u '+%Y-%m-%dT%H:%M:%SZ'`q