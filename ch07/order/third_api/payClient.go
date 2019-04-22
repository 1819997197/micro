package third_api

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	trace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	pb "micro/ch07/order/proto"
)

func GetPayClient() pb.PayService {
	var serviceName = "go.micro.srv.pay"
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379",
		}
	})
	// 初始化服务
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Registry(reg),
		//micro.WrapCall(trace.NewCallWrapper()),
		micro.WrapClient(trace.NewClientWrapper()),
	)

	service.Init()

	return pb.NewPayService(serviceName, service.Client())
}
