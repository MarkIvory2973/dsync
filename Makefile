GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
ifeq ($(GOOS),windows)
GOOUT := dsync_$(GOOS)_$(GOARCH).exe
else
GOOUT := dsync_$(GOOS)_$(GOARCH)
endif

NFPM := nfpm
NFPMFLAGS := --packager deb

# Install dependencies
.PHONY: install
install:
	$(MAKE) -C src install

# Test units
.PHONY: test
test:
	$(MAKE) -C src test

# Build binaries and packages
dist/output:
	$(MAKE) -C src build
	mkdir -p dist
	mv src/output dist

.PHONY: build
build: dist/output

# Build packages
.PHONY: package
package: dist/output
ifeq ($(GOOS),linux)
	$(NFPM) pkg $(NFPMFLAGS) --target dist
endif
	mv dist/output dist/$(GOOUT)

# Clean files
.PHONY: clean
clean:
	$(MAKE) -C src clean

	$(RM) -r dist
