package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/bind", _shouldbind)
	r.POST("/bindjson", _shouldbindjson)
	r.Run(":80")
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func _shouldbind(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func _shouldbindjson(c *gin.Context) {
	var user Login
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	c.JSON(http.StatusOK, user)
}
