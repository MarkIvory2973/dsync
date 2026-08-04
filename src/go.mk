GO := go
GOFLAGS := -trimpath -ldflags="-s -w"

CGO_ENABLED := 0

GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
