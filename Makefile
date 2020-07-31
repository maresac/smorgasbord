SHELL := bash

# Directory, where all required tools are located (absolute path required)
TOOLS_DIR ?= $(shell cd tools && pwd)
# Assets for dex to successfully run tests
DEX_WEB_DIR ?= $(shell cd web && pwd)

# Prerequisite tools
GO ?= go
DOCKER ?= docker

# Tools managed by this project
GINKGO ?= $(TOOLS_DIR)/ginkgo
LINTER ?= $(TOOLS_DIR)/golangci-lint
GOVERALLS ?= $(TOOLS_DIR)/goveralls
GOVER ?= $(TOOLS_DIR)/gover

.EXPORT_ALL_VARIABLES:
.PHONY: test lint fmt vet

test: fmt vet $(GINKGO)
	$(GINKGO) -r -v -cover pkg -- -dex-web-dir=$(DEX_WEB_DIR) $(TEST_FLAGS)


test-%: fmt vet $(GINKGO)
	$(GINKGO) -r -v -cover pkg/$* -- -dex-web-dir=$(DEX_WEB_DIR) $(TEST_FLAGS)

# First run gover to merge the coverprofiles and upload to coveralls
coverage: $(GOVERALLS) $(GOVER)
	$(GOVER)
	$(GOVERALLS) -coverprofile=gover.coverprofile -service=travis-ci -repotoken $(COVERALLS_TOKEN)

lint: $(LINTER)
	$(GO) mod verify
	$(LINTER) run -v --no-config --deadline=5m

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

tools: $(TOOLS_DIR)/ginkgo $(TOOLS_DIR)/golangci-lint $(TOOLS_DIR)/goveralls $(TOOLS_DIR)/gover

$(TOOLS_DIR)/ginkgo:
	$(shell $(TOOLS_DIR)/goget-wrapper github.com/onsi/ginkgo/ginkgo@v1.12.0)

$(TOOLS_DIR)/golangci-lint:
	$(shell curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(TOOLS_DIR) v1.25.0)

$(TOOLS_DIR)/goveralls:
	$(shell $(TOOLS_DIR)/goget-wrapper github.com/mattn/goveralls@v0.0.5)

$(TOOLS_DIR)/gover:
	$(shell $(TOOLS_DIR)/goget-wrapper github.com/modocache/gover)
