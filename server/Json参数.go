package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserId int64 `json:"userId"`
	OrgId  int64 `json:"orgId"`
}

func main() {
	router := gin.Default()

	//query + post 的form 表单

	router.POST("/json", func(c *gin.Context) {

		var user User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
			return
		} else {
			fmt.Printf("user的结构%+v", user)
		}

	})
	router.Run(":8080")
}
