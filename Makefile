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
$(info Checking go environment...)
ifndef GOPATH
	$(error GOPATH is undefined)
endif

$(info Checking required dependencies for build...)
EXECUTABLES = git go dep
K := $(foreach exec,$(EXECUTABLES),\
	$(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

# Install dependencies
install-project-libraries:
	$(info Ensure dependencies are met and download missing ones...)
	@dep ensure

# Test
test:
	$(info Build and run tests...)
	@go test -v -count=1 -timeout 120s github.com/torre76/pgbpasswd/encrypt &>/dev/null

# Build
build: check-env install-project-libraries test
	$(info Build final command to "build/pgbpasswd"...)
	@go build $(LDFLAGS) -o build/pgbpasswd

# Clean
clean:
	@go clean
	rm -rf build vendor
