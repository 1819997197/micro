进入go-kit-order目录
protoc --go_out=plugins=grpc:. ./proto/order.proto

运行服务端
go build -o order  -ldflags '-w -s'
./order

运行客户端
go run cli/main.go

服务端会打印出传入的参数，客户端会输出服务端返回的信息