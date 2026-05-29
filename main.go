package main

import (
	"example.com/dad_midia/db"
	"example.com/dad_midia/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.Use(cors.Default()) // default = todas as origens permitidas
	routes.RegisterRoutes(server)

	server.Run(":8080") // 127.0.0.1:8080
}
