package main

import (
	"github.com/Megidy/TaskManagmentSystem/pkj/config"
	"github.com/Megidy/TaskManagmentSystem/pkj/routes"
	"github.com/Megidy/TaskManagmentSystem/pkj/utils"
	"github.com/gin-gonic/gin"
)

// to do:
// 1 : fix bug with format of time in ToDone and Created
// 2 : instead of returning nil when no value from db ,return err
// 3 : create new file with all the types beacuse there are too musch of them
func init() {
	utils.LoadEnv()
}
func main() {

	router := gin.Default()
	routes.InitRoutes(router)
	router.Run(":8080")
	config.Connect()

}
