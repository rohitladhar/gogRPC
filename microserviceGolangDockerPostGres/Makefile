.PHONY: proto/gen

PROTO_FILES := $(shell find . -name '*.proto')

proto/gen:
	protoc \
	  --go_out=./ \
	  --go_opt=paths=source_relative \
	  --go-grpc_out=./ \
	  --go-grpc_opt=paths=source_relative $(PROTO_FILES)
