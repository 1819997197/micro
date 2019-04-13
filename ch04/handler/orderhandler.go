package handler

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"micro/ch04/conn"
	pb "micro/ch04/proto"
)

type Order struct {}

func InitOrder() *Order {
	return new(Order)
}

func (s *Order) GetOrderInfo(ctx context.Context, req *pb.GetOrderInfoReq, rsp *pb.GetOrderInfoRes) error {
	fmt.Println("received: ", req.Id)
	//1.校验参数(略)

	//2.redis
	c := conn.RedisDB.Get()
	defer c.Close()

	_, err := c.Do("SET", "keyWord", "will")
	if err != nil {
		fmt.Println("redisConn.Do err:", err)
		return err
	}

	keyword, _ := redis.String(c.Do("GET", "keyWord"))
	rsp.Msg = keyword
	return nil
}
