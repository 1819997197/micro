package handler

import (
	"context"
	"fmt"
	pb "micro/ch03/proto"
	"micro/ch03/service"
)

type Order struct {
	OrderService *service.OrderService
}

func InitOrder() *Order {
	o := new(Order)
	o.OrderService = service.InitOrderService()
	return o
}

func (s *Order) GetOrderInfo(ctx context.Context, req *pb.GetOrderInfoReq, rsp *pb.GetOrderInfoRes) error {
	fmt.Println("received: ", req.Id)
	//1.校验参数(略)

	//2.调用service方法
	msg, err := s.OrderService.GetOrderById(req.Id)
	if err != nil {
		fmt.Println("s.OrderService.GetOrderById err:", err)
		return err
	}

	rsp.Msg = msg
	return nil
}
