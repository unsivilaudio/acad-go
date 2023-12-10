package main

import (
	"demoapi/db"
	"demoapi/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("RestAPI started!")
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
