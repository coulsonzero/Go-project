package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	// 禁用控制台颜色(仅记录到文件时)
	// gin.DisableConsoleColor()
	// 强制日志颜色化(需要记录到控制台时)
	// gin.ForceConsoleColor()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
