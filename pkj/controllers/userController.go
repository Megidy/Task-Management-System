package controllers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Megidy/TaskManagmentSystem/pkj/models"
	"github.com/Megidy/TaskManagmentSystem/pkj/types"
	"github.com/Megidy/TaskManagmentSystem/pkj/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SignUp(c *gin.Context) {
	var UserSignUpRequest types.UserAuth
	err := c.ShouldBindBodyWithJSON(&UserSignUpRequest)
	if err != nil {
		utils.HandleError(c, err, "failed to read body ", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(UserSignUpRequest.Password)
	if err != nil {
		utils.HandleError(c, err, "failed to hash password ", http.StatusBadRequest)
		return
	}

	ok, err := models.IsSignedUp(UserSignUpRequest.Username)
	if err != nil {
		utils.HandleError(c, err, "user is already signed up", http.StatusBadRequest)
		return
	}

	if ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, "user is already signed up")
		return
	}

	err = models.CreateUser(UserSignUpRequest.Username, string(hashedPassword))
	if err != nil {
		utils.HandleError(c, err, "failed to create new user", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "you signed up ",
	})
}
func LogIn(c *gin.Context) {
	utils.LoadEnv()
	var UserLogInRequest types.UserAuth

	err := c.ShouldBindBodyWithJSON(&UserLogInRequest)
	if err != nil {
		utils.HandleError(c, err, "failed to read body", http.StatusBadRequest)
	}

	//rework
	ok, err := models.IsSignedUp(UserLogInRequest.Username)
	if err != nil {
		utils.HandleError(c, err, "failed to query row", http.StatusInternalServerError)
	}

	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, "no user found ")
	}

	user, err := models.GetUser(UserLogInRequest.Username)
	if err != nil {
		utils.HandleError(c, err, "failed to get user data ", http.StatusInternalServerError)
	}

	err = utils.UnHashPassword(user.Password, UserLogInRequest.Password)
	if err != nil {
		utils.HandleError(c, err, "failed to unhash password ", http.StatusBadRequest)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 10).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		utils.HandleError(c, err, "failed to create token ", http.StatusInternalServerError)

	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*10, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"succesfull": "you succesfully logged in ",
	})

}

func DeleteUser(c *gin.Context) {
	id := c.Param("userId")
	userId, err := strconv.Atoi(id)
	if err != nil {
		utils.HandleError(c, err, "failed to converte to int ", http.StatusBadRequest)
		return
	}

	err = models.DeleteUser(userId)
	if err != nil {
		utils.HandleError(c, err, "failed to delete user from db ", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "deleted user ",
	})

}
