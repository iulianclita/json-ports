# Install all development tools to the `bin` directory.
export GOBIN=$(CURDIR)/bin

# Default to the system 'go'.
GO?=$(shell which go)

.PHONY: setup
setup: ## Set up local linting tool
	mkdir -p $(GOBIN)
	cd $(GOBIN)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.55.2

.PHONY: clean
clean: ## Remove all build artifacts.
	rm -rf $(GOBIN)

.PHONY: lint
lint: ## Lint the source code.
	$(GOBIN)/golangci-lint run --config $(shell pwd)/build/.golangci.yml --verbose ./...

.PHONY: tests
tests: ## Run all tests
	$(GO) test -v -race ./...

.PHONY: up
up: ## start docker compose
	docker compose up -d

.PHONY: down
down: ## stop docker compose
	docker compose down  
