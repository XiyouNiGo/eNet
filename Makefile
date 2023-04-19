PROJ_NAME = enet
PROM_STACK_NS ?= monitoring
PROM_PORT ?= 9090
GRAF_PORT ?= 9091
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

.PHONY: deploy
deploy: deploy-prom-stack

.PHONY: deploy-prom-stack
deploy-prom-stack:
	@helm upgrade -f ./hack/charts/kube-prometheus-stack/values.yaml --install enet kube-prometheus-stack --namespace ${PROM_STACK_NS} --create-namespace --repo https://prometheus-community.github.io/helm-charts; \
	kubectl -n ${PROM_STACK_NS} expose service ${PROJ_NAME}-grafana --type=NodePort --port=${GRAF_PORT} --target-port=3000 --name=${PROJ_NAME}-grafana-nodeport; \

.PHONY: undeploy
undeploy: undeploy-prom-stack
	
.PHONY: undeploy-prom-stack
undeploy-prom-stack:
	@helm delete ${PROJ_NAME} --namespace ${PROM_STACK_NS}; \
	kubectl -n ${PROM_STACK_NS} delete svc ${PROJ_NAME}-grafana-nodeport; \

.PHONY: create-cluster
create-cluster:
	@kind create cluster --config ./hack/cluster_config.yml

.PHONY: delete-cluster
delete-cluster:
	@kind delete clusters cluster-for-enet

.PHONY: port-forward
port-forward:
	@killall kubectl &>/dev/null; \
  (kubectl -n ${PROM_STACK_NS} port-forward services/${PROJ_NAME}-kube-prometheus-stack-prometheus 8080:${PROM_PORT} > /dev/null &); \
  (kubectl -n ${PROM_STACK_NS} port-forward services/${PROJ_NAME}-grafana-nodeport 8081:${GRAF_PORT} > /dev/null &); \
  echo "Started port-forward commands"; \
  echo "localhost:8080 - prometheus"; \
  echo "localhost:8081 - grafana"; \