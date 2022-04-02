GIT_COMMIT_ID_SHORT := $(shell git rev-parse --short HEAD)
ifneq ($(shell git status --porcelain),)
       GIT_COMMIT_ID_SHORT := $(GIT_COMMIT_ID_SHORT)-dirty
endif

GIT_BRANCH_NAME := $(shell git rev-parse --abbrev-ref HEAD)

BUILD_TIME = `date -u '+%Y-%m-%dT%H:%M:%SZ'`