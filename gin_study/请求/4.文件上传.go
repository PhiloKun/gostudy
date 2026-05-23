package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("upload", func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		println(fileHeader.Filename) //文件名
		println(fileHeader.Size)     //文件大小，单位是字节
		//file, _ := fileHeader.Open()
		//bytes, _ := io.ReadAll(file)
		//os.WriteFile("./"+fileHeader.Filename, bytes, 0666)
		c.SaveUploadedFile(fileHeader, "./upload/"+fileHeader.Filename)
	})

	fmt.Println("app run at http://127.0.0.1:8080")
	r.Run(":8080")
}
