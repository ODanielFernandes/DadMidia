package routes

import (
	"example.com/dad_midia/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/midias", getMidias)
	authenticated.GET("/midias/:id", getMidia)
	authenticated.POST("/midias", createMidia)
	authenticated.PUT("/midias/:id", updateMidia)
	authenticated.DELETE("/midias/:id", deleteMidia)
}
