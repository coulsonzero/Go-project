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

	// curl -X GET http://localhost:8080/users
	r.GET("/users", controllers.GetAllUsers)
	//  curl -X GET http://localhost:8080/users/2
	r.GET("/users/:id", controllers.GetUserById)
	// curl -X POST -H 'Content-Type: application/json' -d '{"name":"john kale", "email": "john@gmail.com"}' http://localhost:8080/users
	r.POST("/users", controllers.CreateUser)
	// curl -X PUT -H 'Content-Type: application/json' -d '{"email": "john@163.com"}' http://localhost:8080/users/17
	r.PUT("/users/:id", controllers.UpdateUser)
	// curl -X DELETE http://localhost:8080/users/17
	r.DELETE("/users/:id", controllers.DeleteUser)

	return r
}
