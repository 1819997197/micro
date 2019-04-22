package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"micro/ch07/api/cli"
	pb "micro/ch07/api/proto"
	"net/http"
	"sync"
)

//首页
func IndexApi(c *gin.Context) {
	var ctx context.Context
	var ok bool
	if ctx, ok = c.Keys["ctx"].(context.Context); !ok {
		ctx = context.Background()
	}
	span := opentracing.SpanFromContext(ctx) //span: SetTag、LogEvent操作可去除

	var id int64 = 1
	var err error
	var order *pb.GetOrderInfoRes
	var user *pb.GetUserInfoRes

	var wg sync.WaitGroup
	wg.Add(2)

	//调用order服务
	go func(id int64) {
		defer wg.Done()
		span.SetTag("service1", "go.micro.srv.order")
		span.LogEvent("call cli.GetOrderClient().GetOrderInfo")
		order, err = cli.GetOrderClient().GetOrderInfo(ctx, &pb.GetOrderInfoReq{Id: id})
		if err != nil {
			fmt.Println("cli.GetOrderClient().GetOrderInfo err:", err)
		}
	}(id)

	//调用user服务
	go func(id int64) {
		defer wg.Done()
		span.SetTag("service2", "go.micro.srv.user")
		span.LogEvent("call cli.GetUserClient().GetUserInfo")
		user, err = cli.GetUserClient().GetUserInfo(ctx, &pb.GetUserInfoReq{Id: 1})
		if err != nil {
			fmt.Println("cli.GetUserClient().GetUserInfo err:", err)
		}
	}(id)

	// goroutine产生的err需要单独处理
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"status": true, "order": order, "user": user})
}
