package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"will/micro/go-kit-order/proto"
)

func main() {
	serviceAddress := "127.0.0.1:50052"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()

	orderClient := order.NewOrderClient(conn)
	res, _ := orderClient.CreateOrder(context.Background(), &order.CreateOrderParam{Msg:"create simple order"})
	fmt.Println(res)
}
