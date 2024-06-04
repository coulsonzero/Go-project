package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	r.GET("/string", _string)
	r.GET("/struct", _struct)
	r.GET("/map", _map)
	r.GET("/json", _json)
	r.GET("xml", _xml)
	r.GET("yaml", _yaml)
	r.LoadHTMLGlob("templates/*")
	r.GET("html", _html)
	r.GET("html", _html)
	r.StaticFile("assets/dash", "assets/dash.png")
	r.GET("/redirect", _redirect)
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}



func _string(c *gin.Context) {
	c.String(200, "hello, gin!")
}

func _struct(c *gin.Context) {
	type User struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
	}

	user := User{"shville", 21}
	c.JSON(200, user)
}

func _map(c *gin.Context) {
	userMap := map[string]string{
		"username": "coulson",
		"age": "23",
	}

	c.JSON(200, userMap)
}

func _json(c *gin.Context) {
	c.JSON(200, gin.H{
		"username": "coulson",
	})
}

func _xml(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK, "data": gin.H{"username": "shville"}})
}

func _yaml(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK, "data": gin.H{"username": "shville"}})
}

func _html(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"username": "shville"})
}

func _redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")		// 302
}


