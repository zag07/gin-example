GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
API_PROTO_FILES=$(shell find api -name *.proto)
CONFIGS_PROTO_FILES=$(shell find internal/conf -name *.proto)

.PHONY: api
# generate api code
api:
	protoc --proto_path=. \
           --proto_path=$$GOPATH/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.1 \
		   --proto_path=$$GOPATH/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20210816210408-a6c8dc7f7220 \
           --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
           $(API_PROTO_FILES)
#           --go-http_out=paths=source_relative:. \  暂时移除这一个

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
		   --go_out=paths=source_relative:. \
		   $(CONFIGS_PROTO_FILES)

.PHONY: wire
wire:
	wire ./cmd/example/...