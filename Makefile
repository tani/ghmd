GO ?= go
TINYGO ?= tinygo
OUT_DIR ?= build
ARCH ?= $(shell $(GO) env GOHOSTARCH)

CMARK_CGO_CFLAGS := -I$(CURDIR)/third_party/cmark-gfm/src \
	-I$(CURDIR)/third_party/cmark-gfm/extensions \
	-I$(CURDIR)/third_party/cmark-gfm/generated

.PHONY: all go-native tinygo-native tinygo-wasi clean

all: go-native tinygo-native tinygo-wasi

go-native:
	mkdir -p "$(OUT_DIR)/go/$(ARCH)"
	CGO_ENABLED=1 CGO_CFLAGS="$(CMARK_CGO_CFLAGS) $(CGO_CFLAGS)" \
		$(GO) build -ldflags='-s -w' -o "$(OUT_DIR)/go/$(ARCH)/ghmd" ./cmd/ghmd

tinygo-native:
	mkdir -p "$(OUT_DIR)/tinygo/$(ARCH)"
	CGO_CFLAGS="$(CMARK_CGO_CFLAGS) $(CGO_CFLAGS)" \
		$(TINYGO) build -no-debug -o "$(OUT_DIR)/tinygo/$(ARCH)/ghmd" ./cmd/ghmd

tinygo-wasi:
	mkdir -p "$(OUT_DIR)/tinygo/wasi"
	CGO_CFLAGS="$(CMARK_CGO_CFLAGS) $(CGO_CFLAGS)" \
		$(TINYGO) build -target=wasi -no-debug -o "$(OUT_DIR)/tinygo/wasi/ghmd.wasm" ./cmd/ghmd

clean:
	rm -f "$(OUT_DIR)/go/$(ARCH)/ghmd" "$(OUT_DIR)/tinygo/$(ARCH)/ghmd" \
		"$(OUT_DIR)/tinygo/wasi/ghmd.wasm" "$(OUT_DIR)/ghmd-go-native" \
		"$(OUT_DIR)/ghmd-native" "$(OUT_DIR)/ghmd.wasm"
