package main

import (
	"github.com/gin-gonic/gin"
	"modern-delivery-service/db"
	"modern-delivery-service/routes"
)

func main() {
	db.ConnectDB()
	server := gin.Default()
	
	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
