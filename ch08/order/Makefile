
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/order.proto
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/pay.proto

.PHONY: build
build: proto

	go build -o order main.go plugin.go
