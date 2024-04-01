package main

import (
	"github.com/gin-gonic/gin"
	"test.com/config"
	routes "test.com/routes"
)

func main() {
	service := gin.Default()
	config.Connect()
	routes.Routes(service)
	service.Run()
}
