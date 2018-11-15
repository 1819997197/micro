package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"will/micro/example1/proto"
)

type Say struct {}

func (s *Say) Hello(ctx context.Context, req *model.SayParam, rsp *model.SayResponse) error {
	fmt.Println("received", req.Msg)
	rsp.Header = make(map[string]*model.Pair)
	rsp.Header["name"] = &model.Pair{Key: 1, Values: "abc"}

	rsp.Msg = "hello world"
	rsp.Values = append(rsp.Values, "a", "b")
	rsp.Type = model.RespType_DESCEND

	return nil
}

func (s *Say) Will(ctx context.Context, req *model.Pair, rsp *model.SayResponse) error {
	fmt.Println("received", req.Key, req.Values)

	var head map[string]*model.Pair
	head = make(map[string]*model.Pair)
	head["name"] = &model.Pair{Key:1}
	rsp.Header = head
	rsp.Msg = "hello will"
	rsp.Values = []string{"a", "b", "c"}
	rsp.Type = model.RespType_NONE
	return nil
}

func main() {
	// 我这里用的etcd 做为服务发现，如果使用consul可以去掉
	/*
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://192.168.3.34:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379",
		}
	})

	// 初始化服务
	service := micro.NewService(
		micro.Name("lp.srv.eg1"),
		micro.Registry(reg),
	)
	*/

	// 如果你用的是consul把上面的注释掉用下面的
	// 初始化服务
	service := micro.NewService(
		micro.Name("will.srv.eg1"),
	)
	service.Init()

	// 注册 Handler
	model.RegisterSayHandler(service.Server(), new(Say))

	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
