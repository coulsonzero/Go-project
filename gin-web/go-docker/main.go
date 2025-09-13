package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "Hello Docker!",
		})
	})
	if err := r.Run(":8080"); err != nil {
		log.Fatal("connect port error")
	}
}
