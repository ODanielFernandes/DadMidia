package main

import (
	"example.com/dad_midia/db"
	"example.com/dad_midia/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // 127.0.0.1:8080
}
