package main

import (
	"github.com/gin-gonic/gin"
	"micro/ch09/api/service"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", service.IndexApi)               //首页
	router.GET("/person", service.GetPersonApi)     //获取单条记录 URL: /person?id=1&name=will

	return router
}
