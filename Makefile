SHELL := /bin/bash
ifeq (,$(shell go env GOBIN))
GOBIN = $(shell go env GOPATH)/bin
else
GOBIN = $(shell go env GOBIN)
endif
GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
GO ?= go
CLANG := clang
CFLAGS := -O2 -g -Wall -Werror $(CFLAGS)

.PHONY: all
all: enet

.PHONY: enet
enet: fmt-go fmt-xdp vet gen-xdp
	@$(GO) build -o ./_output/enet ./main.go

.PHONY: fmt-go
fmt-go:
	@$(GO) fmt ./cmd/... ./pkg/...

.PHONY: gen-xdp
gen-xdp: export BPF_CLANG := $(CLANG)
gen-xdp: export BPF_CFLAGS := $(CFLAGS)
gen-xdp:
	@$(GO) generate ./...

.PHONY: fmt-xdp
fmt-xdp:
	@clang-format --style=Google -i ./pkg/xdp/*.c

.PHONY: vet
vet:
	@$(GO) vet ./cmd/... ./pkg/...

.PHONY: test
test:
	@$(GO) test -v ./... -coverprofile ./_output/coverage.out 

.PHONY: vendor
vendor:
	@$(GO) mod tidy
	@$(GO) mod vendor

.PHONY: cover
cover: fmt-go vet test
	@$(GO) tool cover -html=./_output/coverage.out -o ./_output/coverage.html

.PHONY: clean
clean:
	@rm -rf ./_output