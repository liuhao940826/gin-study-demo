package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//默认 router路由
	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		//获取参数
		name := c.Param("name")

		//返回字符串  code int, format string, values ...interface{}
		c.String(http.StatusOK, "hello %s", name)

	})

	//参数在path里面 用 : 变量 不能匹配 /user
	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		//获取参数
		name := c.Param("name")

		//返回字符串  code int, format string, values ...interface{}
		c.String(http.StatusOK, "hello %s", name)

	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	//启动服务指定端口
	router.Run(":8088")

}
