BUILD_DIR	:= build
VPATH			:= $(BUILD_DIR)

CC				:= go build -i -v -mod vendor
GITHASH 	:= $(shell git rev-parse --short HEAD)
GITBRANCH	:= $(shell git rev-parse --abbrev-ref HEAD)
VERSION				:= $(shell git describe --tags --candidates 1 --match '*.*')
DATE			:= $(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ UTC')
DFLAGS		:= -race
CFLAGS		:= -X 'github.com/ovh/erlenmeyer/cmd.githash=$(GITHASH)' \
	-X 'github.com/ovh/erlenmeyer/cmd.date=$(DATE)' \
	-X 'github.com/ovh/erlenmeyer/cmd.gitbranch=$(GITBRANCH)' \
	-X 'github.com/ovh/erlenmeyer/cmd.version=$(VERSION)'
CROSS			:= GOOS=linux GOARCH=amd64

FORMAT_PATHS	:= ./cmd/ ./core ./middlewares ./proto erlenmeyer.go
LINT_PATHS		:= ./ ./cmd/... ./core/... ./middlewares/...

rwildcard	:= $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))

BUILD_FILE	:= erlenmeyer_$(VERSION)
BUILD_DEST	:= $(BUILD_DIR)/$(BUILD_FILE)

.SECONDEXPANSION:
.PHONY: all
all: format lint dist

.PHONY: init
init:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.27.0

.PHONY: dep
dep:
	go mod vendor -v

.PHONY: tidy
tidy:
	go mod tidy -v

.PHONY: clean
clean:
	rm -rf build
	rm -rf vendor

.PHONY: lint
lint: init
	golangci-lint run

.PHONY: format
format:
	gofmt -w -s $(FORMAT_PATHS)

.PHONY: test
test: dep
	go test -cover ./...

.PHONY: dev
dev: tidy format lint build

.PHONY: build
build: dep erlenmeyer.go $$(call rwildcard, ./cmd, *.go) $$(call rwildcard, ./core, *.go) $$(call rwildcard, ./proto, *.go) $$(call rwildcard, ./middlewares, *.go) $$(call rwildcard, *.go)
	$(CC) $(DFLAGS) -ldflags "$(CFLAGS)" -o $(BUILD_DEST) erlenmeyer.go
	rm -f $(BUILD_DIR)/erlenmeyer
	ln -s $(BUILD_FILE) $(BUILD_DIR)/erlenmeyer

.PHONY: release
release: dep erlenmeyer.go $$(call rwildcard, ./cmd, *.go) $$(call rwildcard, ./core, *.go) $$(call rwildcard, ./proto, *.go) $$(call rwildcard, ./middlewares, *.go) $$(call rwildcard, *.go)
	$(CC) -ldflags "-s -w $(CFLAGS)" -o $(BUILD_DEST) erlenmeyer.go
	rm -f $(BUILD_DIR)/erlenmeyer
	ln -s $(BUILD_FILE) $(BUILD_DIR)/erlenmeyer

.PHONY: dist
dist: dep erlenmeyer.go $$(call rwildcard, ./cmd, *.go) $$(call rwildcard, ./core, *.go) $$(call rwildcard, ./proto, *.go) $$(call rwildcard, ./middlewares, *.go) $$(call rwildcard, *.go)
	$(CROSS) $(CC) -ldflags "-s -w $(CFLAGS)" -o $(BUILD_DEST) erlenmeyer.go
	rm -f $(BUILD_DIR)/erlenmeyer
	ln -s $(BUILD_FILE) $(BUILD_DIR)/erlenmeyer
