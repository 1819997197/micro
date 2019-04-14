package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	pb "micro/ch06/proto"
)

func main() {
	/**
	 * 服务器上面使用etcd可以正常运行，但是本地开发用到别的服务，不可能把所有的服务都在本地跑起来
	 * 本地开发，用到别的服务，但是别的服务在本地环境又没有部署，则可以指定ip+端口或者mock
	 * 1.启动服务时，指定端口。eg:go run srv/main.go --server_address=127.0.0.1:8089
	 *   或者直接去etcd里面查看服务的ip以及端口，具体使用说明见https://github.com/1819997197/micro/tree/master/ch00
	 * 2.调用服务时，指明服务的ip+端口。 eg:c.Call(context.Background(), req, nil, WithAddress(fmt.Sprintf("%s:%d", address, port)))
	 */
	service := micro.NewService(
		micro.Name("go.micro.srv.order"),
	)
	service.Init()

	orderClent := pb.NewOrderService("go.micro.srv.order", service.Client())

	rsp, err1 := orderClent.GetOrderInfo(context.Background(), &pb.GetOrderInfoReq{Id: 1}, client.WithAddress(fmt.Sprintf("%s:%d", "127.0.0.1", 8089)))
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(rsp.Msg)
}
