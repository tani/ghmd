GO ?= go

.PHONY: all build test vet gen check-repro check

all: build test vet

build:
	CGO_ENABLED=0 $(GO) build ./...

test:
	CGO_ENABLED=0 $(GO) test -count=1 ./...

vet:
	CGO_ENABLED=0 $(GO) vet -unsafeptr=false -unreachable=false ./...

gen:
	./scripts/gen-native.sh

check-repro: gen
	git diff --exit-code -- internal/cmark-gfm/cmark_generated.go internal/standard/md4c_generated.go

check: all check-repro
