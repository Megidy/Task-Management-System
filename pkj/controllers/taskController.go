package controllers

import (
	"net/http"
	"time"

	"github.com/Megidy/TaskManagmentSystem/pkj/models"
	"github.com/Megidy/TaskManagmentSystem/pkj/utils"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		utils.HandleError(c, nil, "failed to retrieve data about user", http.StatusUnauthorized)
		return
	}

	var NewUsersTaskRequest struct {
		Title       string
		Description string
		Priority    string
		Dependency  int
		ToDone      time.Time //for example "to_done": "2024-09-30 15:00:00"
	}
	err := c.ShouldBindBodyWithJSON(&NewUsersTaskRequest)
	if err != nil {
		utils.HandleError(c, err, "failed to read body", http.StatusBadRequest)
		return
	}

	var task = models.Task{
		Title:       NewUsersTaskRequest.Title,
		Description: NewUsersTaskRequest.Description,
		Priority:    NewUsersTaskRequest.Priority,
		Dependency:  NewUsersTaskRequest.Dependency,
		ToDone:      NewUsersTaskRequest.ToDone,
		UserId:      user.(*models.User).Id,
	}
	err = models.CreateTask(&task)
	if err != nil {
		utils.HandleError(c, err, "failed to create new task", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "created new task ",
	})
}

func GetTask(c *gin.Context) {

}

func GetAllTasks(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		utils.HandleError(c, nil, "failed to retrieve data about user", http.StatusUnauthorized)
		return
	}
	response, err := models.GetAllTasks(user.(*models.User).Id)
	if err != nil {
		utils.HandleError(c, err, "failed to get tasks", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"your current tasks": response,
	})

}

func UpdateTask(c *gin.Context) {

}

func DeleteTask(c *gin.Context) {

}
