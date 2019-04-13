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

//添加
func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.PostForm("last_name") //接收post方式的参数

	msg := fmt.Sprintf("fist_name: %s, lastName: %s", firstName, lastName)
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}

//获取单条记录
func GetPersonApi(c *gin.Context) {
	id := c.Param("id") //接收get方式的参数

	msg := fmt.Sprintf("id: %s", id)
	c.JSON(http.StatusOK, gin.H{"person": msg})
}
