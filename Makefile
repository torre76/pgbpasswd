SHELL := /bin/bash
BUILD := `git rev-parse HEAD`

.DEFAULT_GOAL := build

# Use linker flags to provide build settings to the target
LDFLAGS=-ldflags "-X=main.Build=$(BUILD)"

.PHONY: check-env install-project-libraries test clean

.SILENT: test clean 

# Pre flight checks
check-env:
ifndef GOPATH
	$(error GOPATH is undefined)
endif

EXECUTABLES = git go 
K := $(foreach exec,$(EXECUTABLES),\
	$(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

install-dep:
	@go get -u github.com/golang/dep/cmd/dep 

install-project-libraries:
	@dep ensure

test:
	@go test -v -count=1 -timeout 120s github.com/torre76/pgbpasswd/encrypt &>/dev/null

build: check-env install-dep install-project-libraries test
	@go build $(LDFLAGS) -o build/pgbpasswd

clean:
	@go clean
	rm -rf build vendor
