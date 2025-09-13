package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-todolist/api"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Hello todolist!")
	})

	v1 := r.Group("api/v1")
	{
		v1.GET("/users", api.ShowUsers)
		v1.GET("/users/:id", api.ShowUser)
		v1.POST("/users/add", api.AddUser)
		v1.PUT("/users/:id", api.UpdateUser)
		v1.DELETE("/users/:id", api.DeleteUser)
	}

	return r
}
