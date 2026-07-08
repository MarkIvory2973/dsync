GO := go
CGO_ENABLED := 0

NFPM := nfpm

.PHONY: install
install:
	cd src && $(GO) mod download

.PHONY: test
test:
	cd src && $(GO) test ./...

.PHONY: build
build:
	cd src && $(GO) build -trimpath -ldflags="-s -w"
	
	mkdir -p dist
	mv src/dsync* dist

.PHONY: package
package:
	$(NFPM) pkg --packager deb --target dist

.PHONY: clean
clean:
	rm -rf dist
