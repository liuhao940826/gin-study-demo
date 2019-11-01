package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//query + post 的form 表单

	router.POST("/post", func(c *gin.Context) {
		//url参数
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")

		//post表单
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	router.Run(":8080")
}
