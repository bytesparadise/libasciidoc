.PHONY: install-ginkgo
## Install development tools.
install-ginkgo:
	@go install -v github.com/onsi/ginkgo/ginkgo

.PHONY: install-gover
## Install development tools.
install-gover:
	@go install -v github.com/sozorogami/gover

.PHONY: test
## run all tests excluding fixtures and vendored packages
test: clean generate-optimized install-ginkgo
	@ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --trace --race --compilers=0

COVERPKGS := $(shell go list ./... | grep -v vendor | paste -sd "," -)

.PHONY: test-with-coverage
## run all tests excluding fixtures and vendored packages
test-with-coverage: generate-optimized install-ginkgo install-gover
	@echo $(COVERPKGS)
	@ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --trace --race --compilers=0  --cover -coverpkg $(COVERPKGS)
	@gover . coverage.txt

.PHONY: test-fixtures
## run all fixtures tests
test-fixtures: generate-optimized
	@ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --trace --race --compilers=0 -tags=fixtures --focus=fixtures
