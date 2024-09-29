package controllers

import (
	"net/http"
	"strconv"

	"github.com/Megidy/TaskManagmentSystem/pkj/models"
	"github.com/Megidy/TaskManagmentSystem/pkj/utils"
	"github.com/gin-gonic/gin"
)

func GetUserLogs(c *gin.Context) {
	id := c.Param("userId")
	userId, err := strconv.Atoi(id)
	if err != nil {
		utils.HandleError(c, err, "failed to converte atring ", http.StatusBadRequest)
	}

	ok, err := models.IsSignedUpById(userId)
	if err != nil {
		utils.HandleError(c, err, "failed to get data from db ", http.StatusInternalServerError)
		return
	}
	if !ok {
		utils.HandleError(c, nil, "no user found", http.StatusBadRequest)
		return
	}

	logs, err := models.GetUsersLogs(userId)
	if err != nil {
		utils.HandleError(c, err, "failed to get logs ", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Logs of all users ": logs,
	})
}

func GetAllLogs(c *gin.Context) {
	logs, err := models.GetAllLogs()
	if err != nil {
		utils.HandleError(c, err, "failed to get logs ", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Logs of all users ": logs,
	})
}
