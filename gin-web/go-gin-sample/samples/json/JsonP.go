package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/jsonp", func(c *gin.Context) {
		// data := map[string]interface{}{
		// 	"foo": "bar",
		// }
		// c.JSONP(http.StatusOK, data)
		c.JSONP(http.StatusOK, gin.H{
			"foo": "bar",
		})
	})
	r.Run(":8080")
}

// {"foo":"bar"}
