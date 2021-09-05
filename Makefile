# copy https://github.com/gin-gonic/gin/blob/master/Makefile
GO ?= go
GOFMT ?= gofmt "-s"
GO_VERSION=$(shell $(GO) version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)
PACKAGES ?= $(shell $(GO) list ./...)
VETPACKAGES ?= $(shell $(GO) list ./... | grep -v /examples/)
GOFILES := $(shell find . -name "*.go")
TESTFOLDER := $(shell $(GO) list ./... | grep -E 'gin$$|binding$$|render$$' | grep -v examples)
TESTTAGS ?= ""

.PHONY dev:
dev: # run webserver
	@$(GO) run wiki.go

.PHONY up:
up: # build and start webserver
	@$(GO) build wiki.go
	./wiki