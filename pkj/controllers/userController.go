package controllers

import (
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var UserSignUpRequest struct {
		Username string
		Password string
	}
	c.ShouldBindBodyWithJSON(&UserSignUpRequest)

}
func LogIn(c *gin.Context) {

}
