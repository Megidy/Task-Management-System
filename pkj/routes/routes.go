package routes

import (
	"github.com/Megidy/TaskManagmentSystem/pkj/controllers"
	"github.com/Megidy/TaskManagmentSystem/pkj/middleware"
	"github.com/gin-gonic/gin"
)

var InitRoutes = func(router gin.IRouter) {
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.LogIn)
	router.GET("/task/:taskId", middleware.RequireAuth, controllers.GetSingleTask)
	router.GET("/tasks", middleware.RequireAuth, controllers.GetAllTasks)
	router.POST("/task", middleware.RequireAuth, controllers.CreateTask)
	router.DELETE("/task/:taskId", middleware.RequireAuth, controllers.DeleteTask)
	// router.PUT("/task/", middleware.RequireAuth, controllers.UpdateTask)
	router.PUT("/task/:taskId", middleware.RequireAuth, controllers.UpdateTask)

}
