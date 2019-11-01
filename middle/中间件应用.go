package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	//全局的日志中间件 和全局的panic处理中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//可以任意的添加对每个路由的中间件个数
	r.GET("/benchmark", firstHandler(), seconHandler())

	//授权组
	authorized := r.Group("/auth")

	//仅对授权组使用中间件 authHandler
	authorized.Use(authHandler())
	{
		authorized.POST("/login", func(context *gin.Context) {

			fmt.Println("授权组里面的 login............")

		})
		authorized.POST("/submit", func(context *gin.Context) {

			fmt.Println("授权组里面的 login............")

		})
		authorized.POST("/read", func(context *gin.Context) {

			fmt.Println("授权组里面的 login............")

		})

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", func(context *gin.Context) {

			fmt.Println("测试组的 服务器异步请求............")
		})
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

func firstHandler() gin.HandlerFunc {

	return func(context *gin.Context) {
		fmt.Println("这是我的第一个:first  handlerFunc")
	}
}

func seconHandler() gin.HandlerFunc {

	return func(context *gin.Context) {
		fmt.Println("这是我的第二个:second  handlerFunc")
	}
}

func authHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("这是授权组 handlerFunc")
	}
}
