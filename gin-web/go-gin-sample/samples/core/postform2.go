package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		// types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("password")
		// c.String(200, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
		c.JSON(200, gin.H{
			"success": true,
			"code":    200,
			"message": "请求成功",
			"data": map[string]string{
				"name":     username,
				"password": password,
			},
		})
	})
	r.Run()
}
