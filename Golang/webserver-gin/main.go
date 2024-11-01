package main

import (
	"example.com/webserver/db"
	"example.com/webserver/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoute(server)

	server.Run(":3000")
}
