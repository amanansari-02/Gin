package router

import (
	"Gin/src/api/controller"
	"Gin/src/config"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := NewRouter()
	router.Run(config.Appconfig.GetString("server.port"))
}

func NewRouter() *gin.Engine {
	router := gin.New()
	resource := router.Group("/api")
	{
		resource.POST("/CreateUser", controller.CreateUser)
	}
	return router
}
