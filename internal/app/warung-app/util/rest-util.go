package util

import (
	"github.com/gin-gonic/gin"
)

func SetContextError(c *gin.Context, errNumber int, message string){
	c.JSON(errNumber, gin.H{
		"status":  errNumber,
		"message": message,
	})
}