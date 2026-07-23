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

# Build binaries
output: $(shell find . -type f -iname "*.go") .install
	$(GO) build $(GOFLAGS) -o output
	-$(UPX) $(UPXFLAGS) output

.PHONY: build
build: output

# Test units
.PHONY: test
test:
	$(GO) test ./...

# Clean files
.PHONY: clean
clean:
	$(RM) .install
	$(RM) output