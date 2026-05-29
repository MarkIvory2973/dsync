GO=go

.PHONY: install
install:
	cd src && $(GO) mod tidy

.PHONY: test
test:
	cd src && $(GO) test ./...

.PHONY: build
build: install test
	cd src && $(GO) build
	mkdir dist
	mv src/dsync dist

.PHONY: clean
clean:
	rm -rf dist
