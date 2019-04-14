package main

import (
	"github.com/micro/go-micro"
	"micro/ch06/handler"
	pb "micro/ch06/proto"
	"time"
)

func main() {
	// 初始化服务
	service := micro.NewService(
		micro.Name("go.micro.srv.order"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20), //设置了30秒的TTL生存期，并设置了每15秒一次的重注册
	)
	service.Init()

	// 注册 Handler
	pb.RegisterOrderHandler(service.Server(), handler.InitOrder())

	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
