package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-gin-swagger/docs"
)

func main() {
	r := gin.Default()

	// http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/hello", handleSwagger)

	r.Run()
}

// 测试接口 godoc
// @tags admin
// @router /hello [get]
// @summary 测试接口
// @description 测试swagger-hello接口
// @accept json
// @produce json
// @param name body string true "用户名，建议使用姓名拼音"
// @success 200 {json} {"success":true,"data":{},"msg":"ok"}
// @failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
func handleSwagger(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to Gin !",
	})
}
