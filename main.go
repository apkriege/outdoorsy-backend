package main

import (
	"fmt"

	"github.com/apkriege/outdoorsy-backend/common/db"
	"github.com/apkriege/outdoorsy-backend/server"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Outdoorsy Backend")
	gin.SetMode(gin.ReleaseMode)

	err := db.Init()
	if err != nil {
		panic(err)
	}

	defer db.CloseDBConnection()
	r := server.SetupRouter()
	r.Run(":8080")
}
