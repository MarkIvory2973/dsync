GO := go
CGO_ENABLED := 0

UPX := upx

NFPM := nfpm

.PHONY: install
install:
	$(GO) install github.com/goreleaser/nfpm/v2/cmd/nfpm@latest
	
	cd src && $(GO) mod download

.PHONY: test
test:
	cd src && $(GO) test ./...

.PHONY: build
build:
	cd src && $(GO) build -trimpath -ldflags="-s -w" -o dsync
	-cd src && $(UPX) --best --lzma dsync
	
	mkdir -p dist
	mv src/dsync dist

.PHONY: package
package:
	$(NFPM) pkg --packager deb --target dist

.PHONY: clean
clean:
	rm -rf src/dsync
	rm -rf dist
