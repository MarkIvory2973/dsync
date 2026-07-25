GO := go
GOFLAGS := -trimpath -ldflags="-s -w"
CGO_ENABLED := 0
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

UPX := upx
UPXFLAGS := --best --lzma

# Install dependencies
.install: go.mod go.sum
	$(GO) mod download

	touch $@

.PHONY: install
install: .install

# Test units
.PHONY: test
test:
	$(GO) test ./...

# Build binaries
dsync: .install $(shell find . -type f -iname "*.go")
	$(GO) build $(GOFLAGS) -o dsync
	-$(UPX) $(UPXFLAGS) dsync

.PHONY: build
build: dsync

# Clean files
.PHONY: clean
clean:
	$(RM) .install
	$(RM) dsync