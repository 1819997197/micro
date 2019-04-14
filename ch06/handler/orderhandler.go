package handler

import (
	"context"
	"fmt"
	pb "micro/ch06/proto"
)

type Order struct {}

func InitOrder() *Order {
	return new(Order)
}

func (s *Order) GetOrderInfo(ctx context.Context, req *pb.GetOrderInfoReq, rsp *pb.GetOrderInfoRes) error {
	fmt.Println("received: ", req.Id)
	//1.校验参数(略)

	rsp.Msg = "test"
	return nil
}
