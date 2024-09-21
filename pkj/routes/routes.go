package routes

import (
	"github.com/Megidy/TaskManagmentSystem/pkj/controllers"
	"github.com/gin-gonic/gin"
)

var InitRoutes = func(router gin.IRouter) {
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.LogIn)

}
