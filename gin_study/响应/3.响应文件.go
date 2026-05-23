package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Wenjian(c *gin.Context) {
	c.Header("Content-Type", "application/octet-stream")              //表示文件流，唤起浏览器下载
	c.Header("Content-Disposition", "attachment; filename=3.响应文件.go") //用来指定下载下来的文件名
	c.File("3.响应文件.go")
}

func main() {
	r := gin.Default()
	r.GET("/", Wenjian)
	fmt.Println("app run at http://127.0.0.1:8083")
	r.Run(":8083")

}
