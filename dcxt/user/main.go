package main

import (
	"github.com/micro/go-micro"
	"micro/dcxt/user/handler"
	pb "micro/dcxt/user/proto"
	"time"
)

func main() {
	// 初始化服务
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20), //设置了30秒的TTL生存期，并设置了每15秒一次的重注册
	)
	service.Init()

	// 注册 Handler
	pb.RegisterUserHandler(service.Server(), handler.InitUser())

	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
