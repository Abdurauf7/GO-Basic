package main

import (
	"practice/Rest/db"
	"practice/Rest/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	server := gin.Default()

	routes.EventRoutes(server)
	routes.UsersRoutes(server)

	server.Run(":8090") // localhost:8090
}
