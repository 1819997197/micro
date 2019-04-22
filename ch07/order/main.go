package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	trace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"log"
	pb "micro/ch07/order/proto"
	"micro/ch07/order/third_api"
	"os"
	"time"
)

type Order struct{}

func (s *Order) GetOrderInfo(ctx context.Context, req *pb.GetOrderInfoReq, rsp *pb.GetOrderInfoRes) error {
	fmt.Println("received_order: ", req.Id)
	//1.校验参数(略)

	//2.处理业务
	span := opentracing.SpanFromContext(ctx)
	span.LogEvent("third_api.GetPayClient().GetPayInfo")
	third_api.GetPayClient().GetPayInfo(ctx, &pb.GetPayInfoReq{Id: 1})

	rsp.Msg = "hello will"
	return nil
}

func main() {
	zipkin_addr := "http://localhost:9411/api/v1/spans"
	hostname, _ := os.Hostname()
	InitTracer(zipkin_addr, hostname, "go.micro.srv.order")

	// 初始化服务
	service := micro.NewService(
		micro.Name("go.micro.srv.order"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20), //设置了30秒的TTL生存期，并设置了每15秒一次的重注册
		micro.WrapHandler(trace.NewHandlerWrapper()),
	)
	service.Init()

	// 注册 Handler
	pb.RegisterOrderHandler(service.Server(), new(Order))

	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}

func InitTracer(zipkinURL string, hostPort string, serviceName string) {
	collector, err := zipkin.NewHTTPCollector(zipkinURL)
	if err != nil {
		log.Fatalf("unable to create Zipkin HTTP collector: %v", err)
		return
	}
	recorder := zipkin.NewRecorder(collector, false, hostPort, serviceName)

	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(true),
		zipkin.TraceID128Bit(true),
	)
	if err != nil {
		log.Fatalf("unable to create Zipkin tracer: %v", err)
		return
	}
	opentracing.InitGlobalTracer(tracer)
	return
}
