# ======================================
VERSION_PACKAGE=github.com/usmhk/goserver/pkg/version
ROOT_APCKAGE=github.com/usmhk/goserver
# ======================================
.DEFAULT_GOAL := all

.PHONY: all
all: build
# ======================================
# Includes

include scripts/make-rules/common.mk
include scripts/make-rules/golang.mk

# ========================================
# Usage

define USAGE_OPTIONS

Options:
	DEBUG Whether to generate debug symbols. Default is 0.

endef

export USAGE_OPTIONS


# ======================================================
.PHONY: build
build:
	$(MAKE) go.build

## help: Show this help info.
.PHONY: help
help: Makefile
	@echo -e "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"
