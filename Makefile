.PHONY: all
all: clean check generate build 

VERSION := $(shell git rev-parse --short HEAD)$(shell git diff-index --quiet HEAD || echo '-dirty')
override BUILD_OPTS += -ldflags '-w -X main.Version=$(VERSION)'

BINARY := bin
$(BINARY):
	mkdir -p bin

GOIMPORTS := bin/goimports
$(GOIMPORTS): $(BINARY)
	go build -o $@ golang.org/x/tools/cmd/goimports

GOLINT := bin/golint
$(GOLINT): $(BINARY)
	go build -o $@ golang.org/x/lint/golint

GINKGO := bin/ginkgo
$(GINKGO): $(BINARY)
	go build -o $@ github.com/onsi/ginkgo/ginkgo

MOCKERY_GEN := bin/mockery
$(MOCKERY_GEN): $(BINARY)
	go build -o $@ github.com/vektra/mockery/v2

bin/course-data-api: generate
	go build $(BUILD_OPTS) -o $@ .

.PHONY: clean
clean:
	rm -rf bin/
	find . -type f 'mock_*.go' -exec rm{} +

.PHONY: check
check: fmt lint vet

.PHONY: fmt
fmt: $(GOIMPORTS)
	@echo "Formatting go files"
	find . -type f '*.go' | xargs $(GOIMPORTS) -w 
	go fmt ./...

.PHONY: lint
lint: $(GOLINT)
	$(GOLINT) ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test $(shell go list ./... | grep -v test) -coverprofile cover.out

.PHONY: generate
generate: $(MOCKERY_GEN)
	go generate ./...

.PHONY: build
build: bin/course-data-api

.PHONY: e2e_test
e2e_test: generate $(GINKGO)
	$(GINKGO) -noColor -nodes=10 -timeout=5m -flakeAttempts=5 ./test/...
