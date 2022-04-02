REPORTS_DIR=./tmp/bench/reports
BENCH_COUNT=10

.PHONY: bench
## run the top-level benchmarks
bench: clean generate-optimized
	@mkdir -p $(REPORTS_DIR)
	@go test -tags bench -bench=. -benchmem -count=$(BENCH_COUNT) -run=XXX \
		github.com/bytesparadise/libasciidoc \
		| tee $(REPORTS_DIR)/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).bench
	@echo "generated $(REPORTS_DIR)/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).bench"

.PHONY: bench-diff
## run the top-level benchmarks and compares with results of 'master' and 'v0.7.0'
bench-diff: clean generate-optimized check-git-status
	@mkdir -p $(REPORTS_DIR)
	@go test -tags bench -bench=. -benchmem -count=$(BENCH_COUNT) -run=XXX \
		github.com/bytesparadise/libasciidoc \
		| tee $(REPORTS_DIR)/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).bench
	@echo "generated $(REPORTS_DIR)/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).bench"
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
	@git checkout $(GIT_BRANCH_NAME)
	@echo ""
	@echo "Comparing with 'master' branch"
	@benchstat $(REPORTS_DIR)/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).bench $(REPORTS_DIR)/master.bench
	@echo ""
	@echo "Comparing with 'v0.7.0' tag"
	@benchstat $(REPORTS_DIR)/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).bench $(REPORTS_DIR)/v0.7.0.bench

check-git-status:
ifneq ("$(shell git status --porcelain)","")
	@echo "Repository contains uncommitted changes:"
	@git status --porcelain 
	@exit 1
else
	@echo "Repository has no uncommitted changes"
endif

