.PHONY: bench
## run the benchmarks on the parser
bench: generate-optimized
	@mkdir -p ./tmp/bench/reports
	@go test -cpuprofile=tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).cpu.prof \
		-memprofile tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).mem.prof \
		-bench=. \
		github.com/bytesparadise/libasciidoc \
		-run=XXX



.PHONY: bench-parser
## run the benchmarks on the parser
bench-parser: generate
	@ginkgo -tags bench -focus "real-world doc-based benchmarks" -memprofile=./tmp/bench/bench.memory pkg/parser
	@ginkgo -tags bench -focus "basic stats" pkg/parser

