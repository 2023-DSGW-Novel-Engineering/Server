package main

import (
	"github.com/2023-DSGW-Novel-Engineering/cation-backend/controller"
	"github.com/2023-DSGW-Novel-Engineering/cation-backend/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	engine := gin.Default()

	// userController
	engine.GET("/users/:id", controller.GetUserInfo)

	engine.Run(":3000")
}
