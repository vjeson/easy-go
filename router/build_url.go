package router

import (
	"demo/controller"
	"demo/middleware"
	"github.com/gin-gonic/gin"
)

func register(router *gin.Engine) {

	router.GET("/users", controller.Users)
	router.POST("/user", controller.Store)
	router.PUT("/user/:id", controller.Update)
	router.DELETE("/user/:id", controller.Destroy)


	router.POST("/login", controller.Login)
	index := router.Group("/index", middleware.Jwt())
	{
		index.GET("/hello", controller.Hello)
	}

}