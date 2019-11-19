package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//创建路由器
	router := gin.Default()

	v1 := router.Group("/v1")
	{

		v1.POST("/login", func(context *gin.Context) {

			fmt.Println("我是v1 login...............")

		})
		v1.POST("/submit", func(context *gin.Context) {

			fmt.Println("我是v1 submit...............")
		})
		v1.POST("/read", func(context *gin.Context) {

			fmt.Println("我是v1 read...............")

		})

	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(context *gin.Context) {

			fmt.Println("我是v2 login~~~~~~~~~~~~~~~~")

		})
		v2.POST("/submit", func(context *gin.Context) {

			fmt.Println("我是v2 submit~~~~~~~~~~~~~~~")
		})
		v2.POST("/read", func(context *gin.Context) {

			fmt.Println("我是v2 read.~~~~~~~~~~~~~~~~")

		})
	}

	//启动服务
	router.Run(":8080")

}
