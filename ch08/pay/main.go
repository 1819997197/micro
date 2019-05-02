package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	trace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"log"
	pb "micro/ch08/pay/proto"
	"os"
	"time"
)

type Pay struct{}

func (u *Pay) GetPayInfo(ctx context.Context, req *pb.GetPayInfoReq, rsp *pb.GetPayInfoRes) error {
	fmt.Println("received_user: ", req.Id)
	//1.校验参数(略)

	//2.处理业务

	rsp.Msg = "user"
	rsp.Values = []string{"will", "allen", "martin"}
	return nil
}

func main() {
	zipkin_addr := "http://localhost:9411/api/v1/spans"
	hostname, _ := os.Hostname()
	InitTracer(zipkin_addr, hostname, "go.micro.srv.pay")

	// 初始化服务
	service := micro.NewService(
		micro.Name("go.micro.srv.pay"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20), //设置了30秒的TTL生存期，并设置了每15秒一次的重注册
		micro.WrapHandler(trace.NewHandlerWrapper()),
	)
	service.Init()

	// 注册 Handler
	pb.RegisterPayHandler(service.Server(), new(Pay))

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
