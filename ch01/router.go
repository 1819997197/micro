package main

import (
	"github.com/gin-gonic/gin"
	"micro/ch01/service"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", service.IndexApi)               //首页
	router.POST("/person", service.AddPersonApi)    //新增
	router.GET("/person/:id", service.GetPersonApi) //获取单条记录

	return router
}
