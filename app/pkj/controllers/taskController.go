package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Megidy/TaskManagmentSystem/pkj/models"
	"github.com/Megidy/TaskManagmentSystem/pkj/types"
	"github.com/Megidy/TaskManagmentSystem/pkj/utils"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		utils.HandleError(c, nil, "failed to retrieve data about user", http.StatusUnauthorized)
		return
	}

	var NewUsersTaskRequest types.TaskRequest
	err := c.ShouldBindBodyWithJSON(&NewUsersTaskRequest)
	if err != nil {
		utils.HandleError(c, err, "failed to read body", http.StatusBadRequest)
		return
	}

	var task = types.Task{
		Title:       NewUsersTaskRequest.Title,
		Description: NewUsersTaskRequest.Description,
		Priority:    NewUsersTaskRequest.Priority,
		Dependency:  NewUsersTaskRequest.Dependency,
		ToDone:      NewUsersTaskRequest.ToDone,
		UserId:      user.(*types.User).Id,
	}
	err = models.CreateTask(task)
	if err != nil {
		utils.HandleError(c, err, "failed to create new task", http.StatusInternalServerError)
		return

	}
	var NewLog = types.Log{
		UserId: user.(*types.User).Id,
		TaskId: task.Id,
		Action: "Created new task ",
	}
	err = models.CreateLog(NewLog)
	if err != nil {
		utils.HandleError(c, err, "failed to Create log ", http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "created new task ",
	})
}

func GetSingleTask(c *gin.Context) {

	user, ok := c.Get("user")
	if !ok {
		utils.HandleError(c, nil, "failed to retrieve data about user", http.StatusUnauthorized)
		return
	}
	id := c.Param("taskId")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		utils.HandleError(c, err, "failed to get param", http.StatusNotFound)
		return
	}
	ok, err = models.IsCreated(taskId, user.(*types.User).Id)
	if err != nil {
		utils.HandleError(c, err, "failed to get data from db", http.StatusInternalServerError)
		return
	}
	if !ok {
		utils.HandleError(c, nil, "no task found", http.StatusBadRequest)
		return
	}
	response, err := models.GetSingleTask(user.(*types.User).Id, taskId)
	if err != nil {
		utils.HandleError(c, err, "failed to retrieve data from db ", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"task : ": response,
	})

}

func GetAllTasks(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		utils.HandleError(c, nil, "failed to retrieve data about user", http.StatusUnauthorized)
		return
	}
	response, err := models.GetAllTasks(user.(*types.User).Id)
	if err != nil {
		utils.HandleError(c, err, "failed to get tasks", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"your current tasks": response,
	})

}

func UpdateTask(c *gin.Context) {
	var UpdatedTask types.TaskUpdateRequest

	err := c.ShouldBindJSON(&UpdatedTask)
	log.Println("task in function :", UpdatedTask)
	if err != nil {
		utils.HandleError(c, err, "failed to read body ", http.StatusBadRequest)
		return
	}
	user, ok := c.Get("user")
	if !ok {
		utils.HandleError(c, nil, "failed to retrieve data about user", http.StatusUnauthorized)
		return
	}
	id := c.Param("taskId")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		utils.HandleError(c, err, "failed to get param", http.StatusNotFound)
		return
	}
	ok, err = models.IsCreated(taskId, user.(*types.User).Id)
	if err != nil {
		utils.HandleError(c, err, "failed to get data from db", http.StatusInternalServerError)
		return
	}
	if !ok {
		utils.HandleError(c, nil, "no task found", http.StatusBadRequest)
		return
	}
	err = models.UpdateTask(UpdatedTask, user.(*types.User).Id, taskId)
	if err != nil {
		utils.HandleError(c, err, "failed to update task", http.StatusInternalServerError)
	}
	var NewLog = types.Log{
		UserId: user.(*types.User).Id,
		TaskId: taskId,
		Action: "Updated task ",
	}
	err = models.CreateLog(NewLog)
	if err != nil {
		utils.HandleError(c, err, "failed to Create log ", http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "you updated task",
	})
}

func DeleteTask(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		utils.HandleError(c, nil, "failed to retrieve user data ", http.StatusUnauthorized)
		return
	}
	id := c.Param("taskId")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		utils.HandleError(c, err, "failed to get taskid", http.StatusNotFound)
		return
	}
	ok, err = models.IsCreated(taskId, user.(*types.User).Id)
	if err != nil {
		utils.HandleError(c, err, "failed to get data from db", http.StatusInternalServerError)
		return
	}
	if !ok {
		utils.HandleError(c, nil, "no task found", http.StatusBadRequest)
		return
	}
	err = models.DeleteTask(user.(*types.User).Id, taskId)
	if err != nil {
		utils.HandleError(c, err, "faield to delete task from db", http.StatusInternalServerError)
		return
	}
	var NewLog = types.Log{
		UserId: user.(*types.User).Id,
		TaskId: taskId,
		Action: "Deleted task ",
	}
	err = models.CreateLog(NewLog)
	if err != nil {
		utils.HandleError(c, err, "failed to Create log ", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "deleted task",
	})

}

func SortTasks(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		utils.HandleError(c, nil, "failed to retrieve data about user", http.StatusUnauthorized)
		return
	}

	sortTitle, err := models.GetTasksByTitle(user.(*types.User).Id)
	if err != nil {
		utils.HandleError(c, err, "failed to get tasks by title", http.StatusInternalServerError)
		return
	}
	sortCreatedAt, err := models.GetTasksByCreatedAt(user.(*types.User).Id)
	if err != nil {
		utils.HandleError(c, err, "failed to get tasks by CreatedAt", http.StatusInternalServerError)
		return
	}
	sortToDone, err := models.GetTasksByToDone(user.(*types.User).Id)
	if err != nil {
		utils.HandleError(c, err, "failed to get tasks by CreatedAt", http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sorted by title (a-z)": sortTitle,
		"sorted by created at ": sortCreatedAt,
		"sorted by to done ":    sortToDone,
	})

}
