GO=go

.PHONY: init
init:
	cd src && $(GO) mod tidy

.PHONY: test
test: init
	cd src && $(GO) test ./...

.PHONY: build
build: init
	cd src && $(GO) build
	
	mkdir dist
	mv src/dsync* dist

.PHONY: clean
clean:
	rm -rf dist
