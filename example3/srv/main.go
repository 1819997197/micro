package main

import (
	"github.com/micro/go-micro"
	"will/micro/example3/comm"
	"will/micro/example3/handler"
	"will/micro/example3/proto/rpcapi"
)

func main() {
	// consul 做为服务发现
	// 初始化服务
	service := micro.NewService(
		micro.Name(comm.ServiceName),
	)
	service.Init()

	// 注册 Handler
	rpcapi.RegisterOrderHandler(service.Server(), new(handler.Order))

	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
