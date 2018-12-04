package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"net"
	"will/micro/go-kit-order/proto"
)

//需要实现order.proto文件中定义的所有方法
type OrderServer struct {
	createOrderHandler grpc_transport.Handler
}

//通过grpc调用CreateOrder时，CreateOrder只做数据透传，调用OrderServer中对应的Handler.ServeGRPC转交给go-kit处理
func (o *OrderServer) CreateOrder(ctx context.Context, in *order.CreateOrderParam) (*order.CreateOrderResponse, error) {
	_, rsp, err := o.createOrderHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*order.CreateOrderResponse), err
}

//创建CreateOrder的EndPoint
func makeCreateOrderEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//创建订单，处理业绩逻辑
		rsp := &order.CreateOrderResponse{}
		rsp.Msg = "create order"
		rsp.Values = append(rsp.Values, "simple order", "scan order")
		rsp.Type = order.RespType_NONE

		var head map[string]*order.Pair
		head = make(map[string]*order.Pair)
		head["a"] = &order.Pair{Key:1, Values:"order1"}
		head["b"] = &order.Pair{Key:2, Values:"order2"}
		rsp.Header = head

		req := request.(*order.CreateOrderParam)
		fmt.Println(req)
		return rsp,nil
	}
}

func decodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func main() {
	//包装OrderServer
	orderServer := new(OrderServer)

	//创建createOrder的Handler
	createOrderHandler := grpc_transport.NewServer(
		makeCreateOrderEndpoint(),
		decodeRequest,
		encodeResponse,
	)

	//OrderServer增加go-kit流程的createOrder处理逻辑
	orderServer.createOrderHandler = createOrderHandler

	//启动grpc服务
	serviceAddress := ":50052"
	ls, _ := net.Listen("tcp", serviceAddress)
	gs := grpc.NewServer()
	order.RegisterOrderServer(gs, orderServer)
	gs.Serve(ls)
}
