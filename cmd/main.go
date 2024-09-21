package main

import (
	"github.com/Megidy/TaskManagmentSystem/pkj/config"
	"github.com/Megidy/TaskManagmentSystem/pkj/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitRoutes(router)
	router.Run(":8080")
	config.Connect()

}
