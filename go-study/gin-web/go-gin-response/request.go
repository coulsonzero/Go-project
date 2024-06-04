package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
 * 查询参数?： c.Query()
 *
 */

func main() {
	r := gin.Default()
	r.GET("/query", _query)
	r.GET("/param/:id", _param)
	r.POST("/form", _form)
	r.POST("/raw", _raw)
	r.Run(":80")
}

// 查询参数
// ?id=2&name=john
func _query(c *gin.Context) {
	name := c.Query("name") // 获取查询参数？未使用返回空字符串
	fmt.Println(name)
	c.String(200, "hello gin!")
}

// 动态参数
func _param(c *gin.Context) {
	fmt.Println(c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"status":  "success",
		"id":      id,
	})
}

func _form(c *gin.Context) {
	fmt.Println(c.PostForm("username"))
	// c.DefaultPostForm("age", "27")
	// form, err := c.MultipartForm()		// 接收所有
	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"status":  "success",
		"data":    c.PostForm("username"),
	})
}

func _raw(c *gin.Context) {
	println(c.GetRawData())

	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"status":  "success",
	})
}
