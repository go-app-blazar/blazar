all: build test

RACE ?= 0
export CGO_ENABLED ?= 0
ifeq ($(RACE), 1)
	GO_RACE := -race
	export CGO_ENABLED := 1
else
	GO_RACE :=
endif

ALL_GO_FILES := $(shell find ./ -name '*.go')

current_dir = $(shell pwd)

.PHONY: build
build:
	go build ./...

.PHONY: clean
clean:
	go clean

.PHONY: test
test:
	go vet ./...
	go test -cover $(GO_RACE) -parallel 10 ./...

.PHONY: format
format:
	go fmt ./...

