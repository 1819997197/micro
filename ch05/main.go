package main

import (
	"context"
	proto "micro/ch05/proto"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
)

// All methods of Sub will be executed when
// a message is received
type Sub struct{}

// Method can be of any name
func (s *Sub) Process(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.1] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

// Alternatively a function can be used
func subEv(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.2] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

func main() {
	// create a service
	service := micro.NewService(
		micro.Name("go.micro.srv.pubsub"),
	)
	// parse command line
	service.Init()

	// 所有订阅者都能收到，防重复处理需接口幂等性
	micro.RegisterSubscriber("example.topic.pubsub.1", service.Server(), new(Sub))

	// 所有订阅者只会有一个收到(可多人订阅，但最终只会有一个服务处理)
	micro.RegisterSubscriber("example.topic.pubsub.2", service.Server(), subEv, server.SubscriberQueue("queue.pubsub"))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
