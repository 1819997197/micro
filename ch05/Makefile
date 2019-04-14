
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/pubsub.proto

.PHONY: build
build: proto

	go build -o ch05 main.go plugin.go
