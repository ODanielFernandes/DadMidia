package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/midias", getMidias)
	server.GET("/midias/:id", getMidia)
	server.POST("/midias", createMidia)
	server.PUT("/midias/:id", updateMidia)
	server.DELETE("/midias/:id", deleteMidia)

	/*
	   authenticated := server.Group("/")
	   authenticated.Use(middlewares.Authenticate)
	   authenticated.GET("/midias", getMidias)
	   authenticated.GET("/midias/:id", getMidia)
	   authenticated.POST("/midias", createMidia)
	   authenticated.PUT("/midias/:id", updateMidia)
	   authenticated.DELETE("/midias/:id", deleteMidia)
	*/
}
