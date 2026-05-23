package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Html(c *gin.Context) {
	c.HTML(200, "index.html", map[string]any{
		"title": "ZKZ",
	})
}

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("响应/templates/*")
	r.LoadHTMLFiles("响应/templates/index.html")
	r.GET("/", Html)
	fmt.Println("app run at http://127.0.0.1:8082")
	r.Run(":8082")

}
