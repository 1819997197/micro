package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	pb "micro/ch08/user/proto"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379",
		}
	})
	// 初始化服务
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Registry(reg),
	)
	service.Init()

	orderClent := pb.NewUserService("go.micro.srv.user", service.Client())

	//调用Hello方法
	rsp, err1 := orderClent.GetUserInfo(context.Background(), &pb.GetUserInfoReq{Id: 1})
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(rsp)
}
