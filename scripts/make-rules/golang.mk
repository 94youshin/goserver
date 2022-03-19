# ========================================
# Makefile helper functions for golang
#

GO := go
GO_LDFLAGS += -X $(VERSION_PACKAGE).GitVersion=$(VERSION) \
		-X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
		-X $(VERSION_PACKAGE).GitTreeState=$(GIT_TREE_STATE) \
		-X $(VERSION_PACKAGE).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

GO_BUILD_FLAGS += -ldflags "$(GO_LDFLAGS)"

.PHONY: go.build
go.build:
	@echo "=================> Building binary"
	@CGO_ENABLE=0 $(GO) build $(GO_BUILD_FLAGS) -o goserver cmd/apiserver/apiserver.go
