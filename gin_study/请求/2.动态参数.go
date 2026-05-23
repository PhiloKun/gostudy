package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"name": name,
		})
	})
	r.GET("/:name/:id", func(c *gin.Context) {
		name := c.Param("name")
		id := c.Param("id")
		c.JSON(200, gin.H{
			"name": name,
			"id":   id,
		})
	})
	fmt.Println("app run at http://127.0.0.1:8080")
	r.Run(":8080")
}
