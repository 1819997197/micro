package main

import (
	"fmt"
	go_config "github.com/micro/go-config"
	"github.com/micro/go-config/source/etcd"
	"github.com/micro/go-micro"
	"micro/ch04/config"
	"micro/ch04/conn"
	"micro/ch04/handler"
	pb "micro/ch04/proto"
	"time"
)

func main() {
	err := loadRedis()
	if err != nil {
		fmt.Println("loadRedis err:", err)
		return
	}

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

func loadRedis() error {
	etcdSource := etcd.NewSource(
		etcd.WithAddress("127.0.0.1:2379"),
		etcd.WithPrefix("/micro/config"),
		etcd.StripPrefix(true),
	)
	conf := go_config.NewConfig()
	if err := conf.Load(etcdSource); err != nil {
		fmt.Println("conf.Load err:", err)
		return err
	}

	var redisConfig config.RedisConfig
	if err := conf.Scan(&redisConfig); err != nil {
		fmt.Println("conf.Scan:", err)
		return err
	}

	conn.InitRedis(redisConfig)
	return nil
}
