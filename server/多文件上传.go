package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	//使用默认的gin

	router := gin.Default()

	router.POST("/mutiUpload", func(c *gin.Context) {

		form, _ := c.MultipartForm()

		files := form.File["files"]

		for _, file := range files {
			log.Println(file.Filename)
		}

		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))

	})

	router.Run(":8080")

}
