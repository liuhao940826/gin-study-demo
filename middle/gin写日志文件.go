package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//禁用控制台颜色
	//gin.DisableConsoleColor()

	gin.ForceConsoleColor()
	//创建路由
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ping 通了")
	})

	r.Run(":8080")
}
