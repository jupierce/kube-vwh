DBG         ?= 0
#REGISTRY    ?= quay.io/openshift/
VERSION     ?= $(shell git describe --always --abbrev=7)
MUTABLE_TAG ?= latest
IMAGE        = $(REGISTRY)machine-api-operator

ifeq ($(DBG),1)
GOGCFLAGS ?= -gcflags=all="-N -l"
endif

.PHONY: all
all: check build test

.PHONY: check
check: lint fmt vet test ## Run code validations

.PHONY: build
build: kube-vwh ## Build binaries

.PHONY: kube-vwh
kube-vwh:
	./hack/go-build.sh kube-vwh
