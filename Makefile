SHELL=/bin/bash -o pipefail

GO ?= go

COMMIT_NO := $(shell git rev-parse HEAD 2> /dev/null || true)
GIT_COMMIT := $(if $(shell git status --porcelain --untracked-files=no),${COMMIT_NO}-dirty,${COMMIT_NO})
GIT_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
GIT_BRANCH_CLEAN := $(shell echo $(GIT_BRANCH) | sed -e "s/[^[:alnum:]]/-/g")

LDFLAGS := -ldflags '-X github.com/sysdiglabs/kubectl-dig/pkg/version.buildTime=$(shell date +%s) -X github.com/sysdiglabs/kubectl-dig/pkg/version.gitCommit=${GIT_COMMIT}'
TESTPACKAGES := $(shell go list ./...)

kubectl_dig ?= _output/bin/kubectl-dig

.PHONY: build
build: clean ${kubectl_dig}

${kubectl_dig}:
	$(GO) build ${LDFLAGS} -o $@ ./cmd/kubectl-dig

.PHONY: clean
clean:
	rm -Rf _output

.PHONY: test
test:
	$(GO) test -v -race $(TESTPACKAGES)