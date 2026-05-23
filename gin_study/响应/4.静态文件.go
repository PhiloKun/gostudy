package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.StaticFile("/abc", "./static/abc.txt")
	fmt.Println("app run at http://127.0.0.1:8084")
	r.Run(":8084")
}
