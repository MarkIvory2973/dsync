# environments
CGO_ENABLED := 0

GO := go
GO_FLAGS := -trimpath -ldflags="-s -w"

.PHONY: init
init:
	cd src && $(GO) mod tidy

.PHONY: test
test: init
	cd src && $(GO) test ./...

.PHONY: build
build: init
	cd src && $(GO) build $(GO_FLAGS)
	
	mkdir dist
	mv src/dsync* dist

.PHONY: clean
clean:
	rm -rf dist
