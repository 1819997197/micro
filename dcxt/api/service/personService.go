package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"micro/dcxt/api/cli"
	"net/http"
	pb "micro/dcxt/api/proto"
	"context"
	"sync"
)

//首页
func IndexApi(c *gin.Context) {
	var id int64 = 1
	var err error
	var order *pb.GetOrderInfoRes
	var user *pb.GetUserInfoRes

	var wg sync.WaitGroup
	wg.Add(2)

	//调用order服务
	go func(id int64) {
		defer wg.Done()
		order, err = cli.GetOrderClient().GetOrderInfo(context.Background(), &pb.GetOrderInfoReq{Id: id})
		if err != nil {
			fmt.Println("cli.GetOrderClient().GetOrderInfo err:", err)
		}
	}(id)

	//调用user服务
	go func(id int64) {
		defer wg.Done()
		user, err = cli.GetUserClient().GetUserInfo(context.Background(), &pb.GetUserInfoReq{Id:1})
		if err != nil {
			fmt.Println("cli.GetUserClient().GetUserInfo err:", err)
		}
	}(id)

	// goroutine产生的err需要单独处理
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"status": true, "order":order, "user":user})
}
