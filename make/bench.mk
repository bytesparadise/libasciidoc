REPORTS_DIR=./tmp/bench/reports
BENCH_COUNT ?= 10

# Detecting GOPATH and removing trailing "/" if any
GOPATH = $(realpath $(shell go env GOPATH))

.PHONY: bench
## run the top-level benchmarks
bench: clean generate-optimized
	@mkdir -p $(REPORTS_DIR)
	@go test -tags bench -bench=. -benchmem -count=$(BENCH_COUNT) -run=XXX \
		github.com/bytesparadise/libasciidoc \
		| tee $(REPORTS_DIR)/$(GITHUB_SHA)-$(GIT_COMMIT_ID_SHORT).bench
	@echo "generated $(REPORTS_DIR)/$(GITHUB_SHA)-$(GIT_COMMIT_ID_SHORT).bench"

.PHONY: bench-diff
## run the top-level benchmarks and compares with results of 'master' and 'v0.7.0'
bench-diff: clean generate-optimized check-git-status
	@git config advice.detachedHead false
	@mkdir -p $(REPORTS_DIR)
	@go test -tags bench -bench=. -benchmem -count=$(BENCH_COUNT) -run=XXX \
		github.com/bytesparadise/libasciidoc \
		| tee $(REPORTS_DIR)/$(GITHUB_SHA)-$(GIT_COMMIT_ID_SHORT).bench
	@echo "generated $(REPORTS_DIR)/$(GITHUB_SHA)-$(GIT_COMMIT_ID_SHORT).bench"
	@git checkout master
	@go test -tags bench -bench=. -benchmem -count=$(BENCH_COUNT) -run=XXX \
		github.com/bytesparadise/libasciidoc \
		| tee $(REPORTS_DIR)/master.bench
	@echo "generated $(REPORTS_DIR)/master.bench"
	@git checkout v0.7.0
	@go test -tags bench -bench=. -benchmem -count=$(BENCH_COUNT) -run=XXX \
		github.com/bytesparadise/libasciidoc \
		| tee $(REPORTS_DIR)/v0.7.0.bench
	@echo "generated $(REPORTS_DIR)/v0.7.0.bench"
	@git checkout $(GITHUB_SHA)
	@echo "HEAD is now at $(shell git log -1 --format='%h') (expecting $(GITHUB_SHA))"
	@$(GOPATH)/bin/benchstat $(REPORTS_DIR)/$(GITHUB_SHA)-$(GIT_COMMIT_ID_SHORT).bench $(REPORTS_DIR)/master.bench >> $(REPORTS_DIR)/diffs-master.txt
	@$(GOPATH)/bin/benchstat $(REPORTS_DIR)/$(GITHUB_SHA)-$(GIT_COMMIT_ID_SHORT).bench $(REPORTS_DIR)/v0.7.0.bench >> $(REPORTS_DIR)/diffs-latest-release.txt

.PHONY: print-bench-diff-master
print-bench-diff-master:
	@cat $(REPORTS_DIR)/diffs-master.txt

.PHONY: print-bench-diff-latest-release
print-bench-diff-latest-release:
	@cat $(REPORTS_DIR)/diffs-latest-release.txt
	
check-git-status:
ifneq ("$(shell git status --porcelain)","")
	@echo "Repository contains uncommitted changes:"
	@git status --porcelain 
	@exit 1
else
	@echo "Repository has no uncommitted changes"
endif

