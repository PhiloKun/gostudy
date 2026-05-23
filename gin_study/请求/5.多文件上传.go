package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("uploads", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)
			return
		}
		for key, headers := range form.File {
			fmt.Println(key, headers)
			for _, h := range headers {
				c.SaveUploadedFile(h, "./uploads/"+h.Filename)
			}
		}

	})

	fmt.Println("app run at http://127.0.0.1:8080")
	r.Run(":8080")
}
