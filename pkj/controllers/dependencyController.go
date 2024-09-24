package controllers

import (
	"net/http"

	"github.com/Megidy/TaskManagmentSystem/pkj/models"
	"github.com/Megidy/TaskManagmentSystem/pkj/utils"
	"github.com/gin-gonic/gin"
)

func GetAllDependencies(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		utils.HandleError(c, nil, "failed to get user data", http.StatusUnauthorized)
		return
	}
	deps, err := models.GetAllDependencies(user.(*models.User).Id)
	if err != nil {
		utils.HandleError(c, err, "failed to get data from db ", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"your current dependencies among tasks  ": deps,
	})

}
