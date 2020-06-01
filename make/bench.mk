.PHONY: bench
## run the benchmarks on the parser
bench: generate-optimized
	@mkdir -p ./tmp/bench/reports
	@go test -cpuprofile=tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).cpu.prof \
		-memprofile tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).mem.prof \
		-bench=. \
		-benchtime=100x \
		github.com/bytesparadise/libasciidoc \
		-run=XXX
	@echo "generate CPU reports..."
	@go tool pprof -text -output=tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).cpu.txt \
		tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).cpu.prof
ifndef CI
	@go tool pprof -svg -output=tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).cpu.svg \
		tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).cpu.prof
endif
	@echo "generate memory reports"
	@go tool pprof -text -output=tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).mem.txt \
		tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).mem.prof
ifndef CI
	@go tool pprof -svg -output=tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).mem.svg \
		tmp/bench/reports/$(GIT_BRANCH_NAME)-$(GIT_COMMIT_ID_SHORT).mem.prof
endif



.PHONY: bench-parser
## run the benchmarks on the parser
bench-parser: generate
	@ginkgo -tags bench -focus "real-world doc-based benchmarks" -memprofile=./tmp/bench/bench.memory pkg/parser
	@ginkgo -tags bench -focus "basic stats" pkg/parser

