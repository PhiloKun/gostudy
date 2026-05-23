package main

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		bytes, _ := io.ReadAll(c.Request.Body)
		fmt.Println(string(bytes))
	})
}
