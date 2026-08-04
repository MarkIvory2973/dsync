include mk/upx.mk
include mk/output.mk
include mk/nfpm.mk

# Install dependencies
.PHONY: install
install:
	$(MAKE) -C src install

# Test units
.PHONY: test
test:
	$(MAKE) -C src test

# Build binaries
dist/output:
	$(MAKE) -C src build
	mkdir -p dist
	mv src/output dist

.PHONY: build
build: dist/output
	-$(UPX) $(UPXFLAGS) output
ifeq ($(OS),linux)
	-$(NFPM) pkg --packager deb $(NFPMFLAGS)
	-$(NFPM) pkg --packager rpm $(NFPMFLAGS)
endif

	mv dist/output dist/$(OUTPUT)

# Clean files
.PHONY: clean
clean:
	$(MAKE) -C src clean

	$(RM) -r dist
