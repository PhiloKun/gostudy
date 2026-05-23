package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("users", func(c *gin.Context) {
		form := c.PostForm("name")
		postForm, ok := c.GetPostForm("age")
		c.JSON(200, gin.H{
			"name": form,
			"age":  postForm,
			"ok":   ok,
		})

	})

	fmt.Println("app run at http://127.0.0.1:8080")
	r.Run(":8080")
}
