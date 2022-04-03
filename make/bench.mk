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
		| tee $(REPORTS_DIR)/$(REF_NAME)-$(GIT_COMMIT_ID_SHORT).bench
	@echo "generated $(REPORTS_DIR)/$(REF_NAME)-$(GIT_COMMIT_ID_SHORT).bench"

.PHONY: bench-diff
## run the top-level benchmarks and compares with results of 'master' and 'v0.7.0'
bench-diff: clean generate-optimized check-git-status
	@mkdir -p $(REPORTS_DIR)
	@go test -tags bench -bench=. -benchmem -count=$(BENCH_COUNT) -run=XXX \
		github.com/bytesparadise/libasciidoc \
		| tee $(REPORTS_DIR)/$(REF_NAME)-$(GIT_COMMIT_ID_SHORT).bench
	@echo "generated $(REPORTS_DIR)/$(REF_NAME)-$(GIT_COMMIT_ID_SHORT).bench"
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
	@git checkout $(REF_NAME)
	@echo "Comparing with 'master' branch"
	@$(GOPATH)/bin/benchstat $(REPORTS_DIR)/$(REF_NAME)-$(GIT_COMMIT_ID_SHORT).bench $(REPORTS_DIR)/master.bench
	@echo ""
	@echo "Comparing with 'v0.7.0' tag"
	@$(GOPATH)/bin/benchstat $(REPORTS_DIR)/$(REF_NAME)-$(GIT_COMMIT_ID_SHORT).bench $(REPORTS_DIR)/v0.7.0.bench | tee $(REPORTS_DIR)/diffs.txt
	@echo "Reports stored in $(REPORTS_DIR)/diffs.txt"

.PHONY: print-bench-diff
print-bench-diff:
	@cat $(REPORTS_DIR)/diffs.txt
	
check-git-status:
ifneq ("$(shell git status --porcelain)","")
	@echo "Repository contains uncommitted changes:"
	@git status --porcelain 
	@exit 1
else
	@echo "Repository has no uncommitted changes"
endif

