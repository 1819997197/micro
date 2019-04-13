package main

import (
	"fmt"
	go_config "github.com/micro/go-config"
	"github.com/micro/go-micro"
	"micro/ch03/config"
	"micro/ch03/conn"
	"micro/ch03/handler"
	pb "micro/ch03/proto"
	"time"
)

func main() {
	err := loadMysql()
	if err != nil {
		fmt.Println("loadMysql err:", err)
		return
	}
	defer conn.SqlDB.Close()

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

func loadMysql() error {
	if err := go_config.LoadFile("./config.yaml"); err != nil {
		fmt.Println("go_config.LoadFile err:", err)
		return err
	}
	var mysqlConfig config.MysqlConfig
	if err := go_config.Get("mysql").Scan(&mysqlConfig); err != nil {
		fmt.Println("go_config.LoadFile err:", err)
		return err
	}
	return conn.InitMysql(&mysqlConfig)
}
