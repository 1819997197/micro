package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"will/micro/example1/proto"
)

func main() {
	// 我这里用的etcd 做为服务发现，如果使用consul可以去掉
	/*reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://192.168.3.34:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379",
		}
	})

	// 初始化服务
	service := micro.NewService(
		micro.Registry(reg),
	)*/

	// 如果你用的是consul把上面的注释掉用下面的

	// 初始化服务
	service := micro.NewService(
		micro.Name("will.srv.eg1"),
	)

	service.Init()

	sayClent := model.NewSayService("will.srv.eg1", service.Client())

	//调用Hello方法
	rsp, err1 := sayClent.Hello(context.Background(), &model.SayParam{Msg: "hello server"})
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(rsp)

	//调用Will方法
	res, err2 := sayClent.Will(context.Background(), &model.Pair{Key:1, Values:"wuy"})
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(res)
}
