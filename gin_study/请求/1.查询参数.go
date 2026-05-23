package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		value := c.Query("name")
		defaultQuery := c.DefaultQuery("age", "13")
		queryArray := c.QueryArray("key")
		c.JSON(200, gin.H{
			"name":         value,
			"defaultQuery": defaultQuery,
			"queryArray":   queryArray,
		})
	})
	fmt.Println("app run at http://127.0.0.1:8080")
	r.Run(":8080")
}
