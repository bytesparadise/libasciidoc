.PHONY: test
## run all tests excluding fixtures and vendored packages
test: generate-optimized
	@echo $(COVERPKGS)
	@ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --trace --race --compilers=0

.PHONY: test-with-coverage
## run all tests excluding fixtures and vendored packages
test-with-coverage: generate-optimized
	@echo $(COVERPKGS)
	@ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --trace --race --compilers=0  --cover -coverpkg $(COVERPKGS)
	@gover . coverage.txt

.PHONY: test-fixtures
## run all fixtures tests
test-fixtures: generate-optimized
	@ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --trace --race --compilers=2 -tags=fixtures --focus=fixtures

.PHONY: bench-parser
## run the benchmarks on the parser
bench-parser: generate-optimized
	$(eval GIT_BRANCH:=$(shell git rev-parse --abbrev-ref HEAD))
	go test -run="XXX" -bench=. -benchmem -count=10 \
		github.com/bytesparadise/libasciidoc/pkg/parser | \
		tee ./tmp/bench-$(GIT_BRANCH).txt
