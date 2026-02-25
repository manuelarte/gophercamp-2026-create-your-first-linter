default: help

help:
	@echo "Please use 'make <target>' where <target> is one of"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z\._-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.PHONY: help

tidy: ## Run go mod tidy in all directories
	go mod tidy
.PHONY: tidy

build: ## Build
	go build .
.PHONY: tidy

lint: tidy ## Format go code and run the fixer
	golangci-lint custom -v
	./custom-gcl fmt
	./custom-gcl run --fix ./...
.PHONY: lint tidy
