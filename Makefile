export GOOS ?= $(shell go env GOOS)
export GOARCH ?= $(shell go env GOARCH)
ifeq ($(GOOS),windows)
export GOOUT := dsync_$(GOOS)_$(GOARCH).exe
else
export GOOUT := dsync_$(GOOS)_$(GOARCH)
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
.PHONY: build
build:
	mkdir -p dist

	$(MAKE) -C src build
	mv src/$(GOOUT) dist
ifeq ($(GOOS),linux)
	$(NFPM) pkg $(NFPMFLAGS) --target dist
endif

# Clean files
.PHONY: clean
clean:
	$(MAKE) -C src clean

	$(RM) -r dist
