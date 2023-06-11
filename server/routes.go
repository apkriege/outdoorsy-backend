package server

import (
	"github.com/apkriege/outdoorsy-backend/controllers/rentals"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/rentals", rentals.GetRentals)
	router.GET("/rentals/:id", rentals.GetRental)

	return router
}
