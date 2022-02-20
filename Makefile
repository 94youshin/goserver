# ======================================
# Includes

include scripts/make-rules/common.mk

# ========================================
# Usage

define USAGE_OPTIONS

Options:
	DEBUG Whether to generate debug symbols. Default is 0.

endef

export USAGE_OPTIONS

## help: Show this help info.
.PHONY: help
help: Makefile
	@echo -e "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"
