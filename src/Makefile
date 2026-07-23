GO := go
GOFLAGS := -trimpath -ldflags="-s -w"
CGO_ENABLED := 0

UPX := upx
UPXFLAGS := --best --lzma

# Install dependencies
.install: go.mod go.sum
	$(GO) mod download

	touch $@

.PHONY: install
install: .install

# Build binaries
$(GOOUT): $(shell find . -type f -iname "*.go") .install
	$(GO) build $(GOFLAGS) -o $(GOOUT)
	-$(UPX) $(UPXFLAGS) $(GOOUT)

.PHONY: build
build: $(GOOUT)

# Test units
.PHONY: test
test:
	$(GO) test ./...

# Clean files
.PHONY: clean
clean:
	$(RM) .install
	$(RM) $(GOOUT)