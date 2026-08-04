include mk/go.mk

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
output: .install $(shell find . -type f -iname "*.go")
	$(GO) build $(GOFLAGS) -o output

.PHONY: build
build: output

# Clean files
.PHONY: clean
clean:
	$(RM) .install
	$(RM) output