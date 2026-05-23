package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": "200",
		"Data": "Hello World",
		"Msg":  "成功！！！",
	})
}

func main() {
	r := gin.Default()
	r.GET("/index", Index)
	fmt.Println("app run at http://127.0.0.1:8081")
	r.Run(":8081")

}
