GIT_SHA=$(shell git rev-parse --verify HEAD)
GOBUILD_CLI=go build -o ./bin/maas-cli

clean: ## clean build output
	rm -rf bin/*

compile: ## build and place in local bin directory
	${GOBUILD_CLI} ./cmd/maascli

compile-linux: ## build linux version
	GOOS=linux GOARCH=amd64 ${GOBUILD_CLI} -o bin/linux-amd64 ./cmd/maascli

go-lint-checks: ## run linting checks against golang code
	./scripts/verify.sh

go-clean: ## have gofmt and goimports clean up go code
	gofmt -w -s ./pkg
	gofmt -w -s ./pkg/cli/
	gofmt -w -s ./pkg/util/
	gofmt -w -s ./cmd
	gofmt -w -s ./cmd/maascli/
	goimports -w $(shell git ls-files "**/*.go" "*.go" | grep -v -e "vendor" | xargs echo)

.PHONY: install-tools
install-tools: ## install tools needed by go-link-checks
	GOIMPORTS_CMD=$(shell command -v goimports 2> /dev/null)
ifndef GOIMPORTS_CMD
	go get golang.org/x/tools/cmd/goimports
endif

	GOLINT_CMD=$(shell command -v golint 2> /dev/null)
ifndef GOLINT_CMD
	go get github.com/golang/lint/golint
endif

	GOCYCLO_CMD=$(shell command -v gocyclo 2> /dev/null)
ifndef GOLINT_CMD
	go get github.com/fzipp/gocyclo
endif

.PHONY: help
help:  ## Show help messages for make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'
