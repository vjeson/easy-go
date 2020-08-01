package router

import (
	"demo/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	router.Static("/static", "static")

	router.Use(middleware.Cors())

	router.Use(middleware.LoggerToFile())
	//router.Use(middleware.LoggerToStd())


	register(router)

	return router
}


