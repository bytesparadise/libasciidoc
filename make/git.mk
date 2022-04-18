GIT_COMMIT_ID_SHORT := $(shell git rev-parse --short HEAD)
ifneq ($(shell git status --porcelain),)
       GIT_COMMIT_ID_SHORT := $(GIT_COMMIT_ID_SHORT)-dirty
endif

# using the env var defined by GitHub Actions 
# (see https://docs.github.com/en/actions/learn-github-actions/environment-variables#default-environment-variables)
GITHUB_SHA ?= $(shell git rev-parse --abbrev-ref HEAD)

BUILD_TIME = `date -u '+%Y-%m-%dT%H:%M:%SZ'`