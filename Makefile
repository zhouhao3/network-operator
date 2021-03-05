KUSTOMIZE=./tools/kustomize
CONTROLLER_GEN=./tools/controller-gen

# Image URL to use all building/pushing image targets
IMG ?= controller:latest

# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Build manager binary
build: generate
	go build -o bin/manager main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate
	go run ./main.go

# Build the docker image
docker: generate
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o bin/manager-docker main.go
	docker rmi -f ${IMG}
	docker build -f build/Dockerfile -t ${IMG} .
	docker save -o ./bin/${IMG}.tar ${IMG}

# Install CRDs into a cluster
install: manifests
	$(KUSTOMIZE) build config/crd | kubectl apply -f -

# Uninstall CRDs from a cluster
uninstall: manifests
	$(KUSTOMIZE) build config/crd | kubectl delete -f -

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | kubectl apply -f -

# Generate code
generate:
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

# Generate manifests e.g. CRD, RBAC etc.
manifests:
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

# Generate docs
.PHONY: docs
docs:
	./tools/plantuml docs/*/*.plantuml

# Run tests
test: generate gofmt golint govet gosec unit manifests

# Run go fmt against code
gofmt:
	./hack/gofmt.sh

# Run go lint against code
golint:
	./hack/golint.sh

# Run go vet against code
govet:
	go vet ./...

# Run go sec against code
gosec:
	./hack/gosec.sh

# Run go test against code
unit:
	go test ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o coverage.html

# Clean files generated by scripts
clean:
	rm -f ./cover.out
	rm -f ./coverage.html
	rm -f ./bin/*

# Clean up go module settings
mod:
	go mod tidy
	go mod verify
