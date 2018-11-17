package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"will/micro/example3/comm"
	"will/micro/example3/proto/rpcapi"
	"will/micro/example3/proto/model"
)

func main() {
	// 服务发现consul
	// 初始化服务
	service := micro.NewService(
		micro.Name(comm.ServiceName),
	)

	service.Init()

	orderClent := rpcapi.NewOrderService(comm.ServiceName, service.Client())

	//调用Hello方法
	rsp, err1 := orderClent.Hello(context.Background(), &model.SayParam{Msg: "hello server"})
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(rsp)
}
