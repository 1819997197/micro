运行下面的代码生成protobuf和micro代码文件
protoc --proto_path=. --micro_out=. --go_out=. you path/hello.proto

micro api（localhost：8080） - 作为http入口点
api服务（go.micro.api.greeter） - 为面向公众提供服务
后端服务（go.micro.srv.greeter） - 内部范围服务

consul agent -dev
启动服务go.micro.srv.greeter
go run greeter/srv/main.go --server_address=127.0.0.1:9090
启动API服务go.micro.api.greeter
go run examples/greeter/api/api.go (未指定端口)
开始 Micro API
micro api

通过micro API进行HTTP调用
Curl "http://localhost:8080/greeter/will/infos?name=Asim+Aslam"
HTTP path/greeter/will/infos 映射到服务 go.micro.api.greeter 方法Say.Hello