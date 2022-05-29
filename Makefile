prefix?=$(GOPATH)/bin
root?=$(shell pwd)
name?=$(notdir $(root))
srcpath?=$(root)/cmd/$(name)
builddir?=$(root)/bin



goos?=$(shell go env GOOS)
arch?=$(shell go env GOARCH)
ldflags?=-ldflags "-s -w"



mansrc?=$(root)/README.md
ifeq ($(shell uname),Linux)
	mandir?=/usr/local/man/man1
endif
ifeq ($(shell uname),Darwin)
	mandir?=/usr/local/share/man/man1
endif
ifeq ($(OS),Windows_NT)
	ext?=.exe
endif



.PHONY: tidy build man install pre post clean
default: all
all: $(prefix) pre tidy install post



pre: $(shell which go) $($(root)/go.mod)
	go mod tidy
	go clean
	mkdir -p $(builddir)
post:
	rm -rf $(builddir)



build: $(prefix) pre $(builddir)
	GOOS=$(goos) GOARCH=$(arch) go build $(ldflags) -o $(builddir)/$(name)$(ext) $(srcpath)
man: $(shell which pandoc) $(mandir) $(builddir)
	@pandoc $(mansrc) -s -t man -o $(builddir)/$(name).1



install: build man
	mv $(builddir)/$(name)$(ext) $(prefix)/
	mv $(builddir)/$(name).1 $(mandir)/
uninstall:
	rm -f $(prefix)/$(name)$(ext)
	rm -f $(mandir)/$(name).1



test:
	go test ./... -coverprofile=cover.out
	curl -Ls https://coverage.codacy.com/get.sh -o codacy.sh && \
	bash ./codacy.sh report -s --force-coverage-parser go -r cover.out -t ${CODACY_PROJECT_TOKEN}
lint:
	golangci-lint run ./... -c ./.golangci.yml
