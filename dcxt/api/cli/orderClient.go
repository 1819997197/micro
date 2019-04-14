package cli

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	pb "micro/dcxt/api/proto"
)

func GetOrderClient() pb.OrderService {
	var serviceName = "go.micro.srv.order"
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379",
		}
	})
	// 初始化服务
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Registry(reg),
	)

	service.Init()

	return pb.NewOrderService(serviceName, service.Client())
}
