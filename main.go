package main

import (
	"go-socialmedia/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	
	// Serve file upload
	r.Static("/uploads", "./uploads")

	// Setup routes
	routes.SetupRoutes(r)

	r.Run(":8080") // listen on localhost:8080
}
