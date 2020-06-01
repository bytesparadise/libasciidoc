GIT_COMMIT_ID := $(shell git rev-parse HEAD)
ifneq ($(shell git status --porcelain --untracked-files=no),)
       GIT_COMMIT_ID := $(GIT_COMMIT_ID)-dirty
endif

GIT_COMMIT_ID_SHORT := $(shell git rev-parse --short HEAD)
ifneq ($(shell git status --porcelain --untracked-files=no),)
       GIT_COMMIT_ID_SHORT := $(GIT_COMMIT_ID_SHORT)-dirty
endif

GIT_BRANCH_NAME := $(shell git rev-parse --abbrev-ref HEAD)

BUILD_TIME = `date -u '+%Y-%m-%dT%H:%M:%SZ'`