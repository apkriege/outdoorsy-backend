package server

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/users", getRentals)
	// r.GET("/api/users/:id", getUser)
	// r.POST("/api/users", createUser)
	// r.PUT("/api/users/:id", updateUser)
	// r.DELETE("/api/users/:id", deleteUser)

	return r
}

func getRentals(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
