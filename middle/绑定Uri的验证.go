package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		var person Person

		if err := c.ShouldBindUri(&person); err != nil {

			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"name": person.Name, "id": person.ID})
	})
	r.Run(":8080")

}
