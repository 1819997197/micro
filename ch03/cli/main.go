package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	pb "micro/ch03/proto"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379",
		}
	})
	// 初始化服务
	service := micro.NewService(
		micro.Name("go.micro.srv.order"),
		micro.Registry(reg),
	)
	service.Init()

	orderClent := pb.NewOrderService("go.micro.srv.order", service.Client())

	//调用Hello方法
	rsp, err1 := orderClent.GetOrderInfo(context.Background(), &pb.GetOrderInfoReq{Id: 1})
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(rsp.Msg)
}
