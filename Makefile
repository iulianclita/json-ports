# Install all development tools to the `bin` directory.
export GOBIN=$(CURDIR)/bin

# Default to the system 'go'.
GO?=$(shell which go)

$(GOBIN):
	mkdir -p $(GOBIN)

.PHONY: setup-lint
setup: ## Set up local linting tool
	cd $(GOBIN)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.55.2

.PHONY: clean
clean: ## Remove build artifacts.
	rm -rf $(GOBIN)

.PHONY: lint
lint: ## Lint the source code.
	$(GOBIN)/golangci-lint run --config $(shell pwd)/build/.golangci.yml --verbose ./...

.PHONY: tests
tests: ## Run all test (unit + integration)
	$(GO) test -v -race -tags=integration ./...

.PHONY: unit-tests
tests: ## Run only unit tests
	$(GO) test -v -race ./...
