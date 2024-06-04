package router

import (
	"gin-grom-mysql/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// /users
	r.GET("/users", controllers.GetAllUsers)
	// /users/2
	r.GET("/users/:id", controllers.GetUserById)
	// users '{"name":"john kale", "email": "john@gmail.com"}'
	r.POST("/users", controllers.CreateUser)
	// /users/2 '{"email": "john@163.com"}'
	r.PUT("/users/:id", controllers.UpdateUser)
	// /users/2
	r.DELETE("/users/:id", controllers.DeleteUser)

	return r
}
