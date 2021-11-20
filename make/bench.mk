.PHONY: bench
## run the top-level benchmarks
bench: clean generate-optimized
	@mkdir -p ./tmp/bench/reports
	@go test github.com/bytesparadise/libasciidoc -run TestParseBasicDocument
	@go test -bench=. -benchmem -count=10 -run=XXX \
		github.com/bytesparadise/libasciidoc \
		| tee tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).bench
	@echo "generated tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).bench"

.PHONY: bench-smoke
## smoke test the top-level benchmarks
bench-smoke: generate-optimized
	@go test -bench=. -benchmem -benchtime=1x -run=XXX \
		github.com/bytesparadise/libasciidoc

# .PHONY: bench-parser
# ## run the benchmarks on the parser
# bench-parser: generate
# 	@ginkgo -tags bench -focus "real-world doc-based benchmarks" -memprofile=./tmp/bench/bench.memory pkg/parser
# 	@ginkgo -tags bench -focus "basic stats" pkg/parser

