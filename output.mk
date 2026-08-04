GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

ifeq ($(GOOS),windows)
OUTPUT := dsync_$(GOOS)_$(GOARCH).exe
else
OUTPUT := dsync_$(GOOS)_$(GOARCH)
endif