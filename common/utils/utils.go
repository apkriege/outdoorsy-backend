package utils

import (
	"github.com/gin-gonic/gin"
)

func ReturnSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"error": "false",
		"data":  data,
	})
}

func ReturnError(c *gin.Context, err string) {
	c.JSON(500, gin.H{
		"error":   "true",
		"message": err,
	})
}
