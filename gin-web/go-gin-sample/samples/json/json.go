package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"code":    200,
			"message": "请求成功",
		})
	})
	if err := r.Run(":8080"); err != nil {
		log.Fatal("connect to gin error !")
	}
}

// {"code":200,"message":"请求成功","success":true}
