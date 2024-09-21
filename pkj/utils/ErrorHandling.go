package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, message string, statusCode int) {

	log.Println("error: ", err, " details", message)
	c.JSON(statusCode, gin.H{
		"details": message,
		"error":   err,
	})
	c.AbortWithStatus(statusCode)

}
