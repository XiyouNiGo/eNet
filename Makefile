PROJ_NAME = enet
PROM_STACK_NS ?= monitoring
PROM_SERVICE_PORT ?= 9090
GRAF_SERVICE_PORT ?= 9091
PROM_NODEPORT ?= 30000
GRAF_NODEPORT ?= 30001
PROM_TARGET_PORT ?= 9090
GRAF_TARGET_PORT ?= 3000
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
COMMITID = $(shell git rev-parse --short HEAD)

.PHONY: all
all: enet-cli enet-exporter

.PHONY: enet-cli
enet-cli: fmt-go fmt-xdp vet gen-xdp
	@$(GO) build -o ./_output/$(GOOS)/$(GOARCH)/enet ./cmd/enet-cli/main.go

.PHONY: enet-exporter
enet-exporter: fmt-go fmt-xdp vet gen-xdp
	@$(GO) build -o ./_output/$(GOOS)/$(GOARCH)/enet-exporter ./cmd/enet-exporter/main.go

.PHONY: cross
cross:
	@mkdir -p _output; \
	GOOS=linux GOARCH=amd64 make; \
	GOOS=linux GOARCH=arm64 make; \

.PHONY: cross-in-docker
cross-in-docker:
	@mkdir -p _output; \
	GOOS=linux GOARCH=amd64 ./hack/dockerized make; \
	GOOS=linux GOARCH=arm64 ./hack/dockerized make; \

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

.PHONY: cover-in-docker
cover-in-docker:
	@mkdir -p _output; \
	GOOS=linux GOARCH=amd64 ./hack/dockerized make cover; \

.PHONY: build-img
build-img:
	@cp -f ./hack/Dockerfile ./_output; \
	docker buildx build --platform=linux/amd64,linux/arm64 \
		-t github.com/xiyounigo/enet:$(COMMITID) ./_output/; \
	echo github.com/xiyounigo/enet:$(COMMITID) > ./_output/images.txt; \

.PHONY: clean
clean:
	@rm -rf ./_output

.PHONY: clean-in-docker
clean-in-docker:
	@./hack/dockerized make clean

.PHONY: deploy
deploy: deploy-prom-stack

.PHONY: deploy-prom-stack
deploy-prom-stack:
	@helm upgrade -f ./hack/charts/kube-prometheus-stack/values.yaml --install enet kube-prometheus-stack --namespace ${PROM_STACK_NS} --create-namespace --repo https://prometheus-community.github.io/helm-charts; \
	kubectl -n ${PROM_STACK_NS} delete service ${PROJ_NAME}-grafana > /dev/null; \
	kubectl -n ${PROM_STACK_NS} expose deployment ${PROJ_NAME}-grafana --type=NodePort --port=${GRAF_SERVICE_PORT} --target-port=${GRAF_TARGET_PORT} --name=${PROJ_NAME}-grafana --overrides='{ "apiVersion": "v1","spec":{"ports": [{"port":${GRAF_SERVICE_PORT},"protocol":"TCP","targetPort":${GRAF_TARGET_PORT},"nodePort":${GRAF_NODEPORT}}]}}'\

.PHONY: undeploy
undeploy: undeploy-prom-stack
	
.PHONY: undeploy-prom-stack
undeploy-prom-stack:
	@helm delete ${PROJ_NAME} --namespace ${PROM_STACK_NS}; \

.PHONY: create-cluster
create-cluster:
	@kind create cluster --config ./hack/cluster_config.yaml

.PHONY: delete-cluster
delete-cluster:
	@kind delete clusters cluster-for-enet

.PHONY: port-forward
port-forward:
	@killall kubectl &> /dev/null; \
  (kubectl -n ${PROM_STACK_NS} port-forward services/${PROJ_NAME}-kube-prometheus-stack-prometheus 8000:${PROM_SERVICE_PORT} > /dev/null &); \
  (kubectl -n ${PROM_STACK_NS} port-forward services/${PROJ_NAME}-grafana 8001:${GRAF_SERVICE_PORT} > /dev/null &); \
  echo "Started port-forward commands"; \
  echo "localhost:8000 - prometheus"; \
  echo "localhost:8001 - grafana"; \