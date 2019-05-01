package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//首页
func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "it works!")
}

//获取单条记录
func GetPersonApi(c *gin.Context) {
	id := c.Query("id")
	name := c.DefaultQuery("name", "guest")

	msg := fmt.Sprintf("id: %s, name: %s", id, name)
	c.JSON(http.StatusOK, gin.H{"person": msg})
}
