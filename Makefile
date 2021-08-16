GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
CONFIGS_PROTO_FILES=$(shell find internal/conf -name *.proto)

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
			   --go_out=paths=source_relative:. \
			   $(CONFIGS_PROTO_FILES)

.PHONY: wire
wire:
	wire ./cmd/example/...