SHELL := /bin/bash
BUILD := `git rev-parse HEAD`

.DEFAULT_GOAL := build

# Use linker flags to provide build settings to the target
LDFLAGS=-ldflags "-X=main.Build=$(BUILD)"

# Tasks that not produce any files
.PHONY: check-env install-project-libraries test clean

# Tasks that would have not echo on execution
.SILENT: test clean 

# Pre flight checks
check-env:
ifndef GOPATH
	$(error GOPATH is undefined)
endif

EXECUTABLES = git go 
K := $(foreach exec,$(EXECUTABLES),\
	$(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

# Install dependency manager
install-dep:
	@go get -u github.com/golang/dep/cmd/dep 

# Install dependencies
install-project-libraries:
	@dep ensure

# Test
test:
	@go test -v -count=1 -timeout 120s github.com/torre76/pgbpasswd/encrypt &>/dev/null

# Build
build: check-env install-dep install-project-libraries test
	@go build $(LDFLAGS) -o build/pgbpasswd

# Clean
clean:
	@go clean
	rm -rf build vendor
