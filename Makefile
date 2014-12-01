GOPATH := $(shell pwd)
.PHONY: clean test

all:
	@GOPATH=$(GOPATH) go install qtunnel

clean:
	@rm -fr bin pkg

dep:
	@GOPATH=$(GOPATH) go get github.com/mreiferson/go-snappystream

test:
	@GOPATH=$(GOPATH) go test tunnel -timeout 2s
