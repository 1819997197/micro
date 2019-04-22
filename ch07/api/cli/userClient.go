package cli

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	trace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	pb "micro/ch07/api/proto"
)

func GetUserClient() pb.UserService {
	var serviceName = "go.micro.srv.user"
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

	return pb.NewUserService(serviceName, service.Client())
}
