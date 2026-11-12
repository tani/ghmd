#!/bin/sh
set -eu

# Generate the Go implementations from unmodified upstream C sources.
# Run from the repository root. CCGO must point at ccgo's module MAIN.
: "${CCGO_VERSION:=v4.39.0}"

go run "modernc.org/ccgo/v4@${CCGO_VERSION}" --goptr=off \
  --package-name=standard \
  -o internal/standard/md4c_generated.go \
  third_party/md4c/md4c.c \
  third_party/md4c/md4c-html.c \
  third_party/md4c/entity.c \
  third_party/md4c-bridge/bridge.c

go run "modernc.org/ccgo/v4@${CCGO_VERSION}" --goptr=off \
  --package-name=cmarkghmd \
  -Ithird_party/cmark-gfm/src \
  -Ithird_party/cmark-gfm/extensions \
  -Ithird_party/cmark-gfm/generated \
  -o internal/cmark-gfm/cmark_generated.go \
  third_party/cmark-gfm-bridge/sources.c
