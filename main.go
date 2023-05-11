package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isaaclyra132/petshop-go/config"
	"github.com/isaaclyra132/petshop-go/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.AppRoutes(router)
	router.Run(":8030")
}
