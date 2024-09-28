package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Megidy/TaskManagmentSystem/pkj/models"
	"github.com/Megidy/TaskManagmentSystem/pkj/types"
	"github.com/Megidy/TaskManagmentSystem/pkj/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	utils.LoadEnv()
	tokenstring, err := c.Cookie("Authorization")
	if err != nil {
		utils.HandleError(c, err, "failed to get cookie", http.StatusUnauthorized)
	}
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(" Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			utils.HandleError(c, err, "your token expired", http.StatusUnauthorized)
		}
		user, err := models.FindUserById(claims["sub"].(float64))
		if err != nil {
			utils.HandleError(c, err, "failed to retrive user from db", http.StatusInternalServerError)

		}
		c.Set("user", user)
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func RequireAdmin(c *gin.Context) {

	user, ok := c.Get("user")
	if !ok {
		utils.HandleError(c, nil, "failed to retrieve user data", http.StatusNotFound)
		return
	}
	if user.(*types.User).Role == "admin" {
		c.Next()
	} else {
		utils.HandleError(c, nil, "", http.StatusNotFound)
		return
	}

}
