.PHONY: bench
## run the top-level benchmarks
bench: clean generate-optimized
	@mkdir -p ./tmp/bench/reports
	@go test -tags bench -bench=. -benchmem -count=10 -run=XXX \
		github.com/bytesparadise/libasciidoc \
		| tee tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).bench
	@echo "generated tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).bench"
