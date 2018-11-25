grpc微服务实例
protoc --go_out=plugins=grpc:. book.proto

grpc插件指定context路径contextPkgPath = “golang.org/x/net/context”可能会与我们代码中的context冲突
把book.proto中import”golang.org/x/net/context” 更改为 “context”

运行服务
go run src/main.go

运行客户端
go run cli/main.go