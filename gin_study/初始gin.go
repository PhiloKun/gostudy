package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// handler
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func main() {
	//1.初始化
	r := gin.Default()
	//线上，关闭日志
	gin.SetMode(gin.ReleaseMode)
	//2.挂载路由
	r.GET("/index", Index)
	//3.绑定端口运行
	fmt.Println("app run at http://127.0.0.1:8080")
	r.Run(":8080") //等价于r.Run("0.0.0.0:8080")

}
