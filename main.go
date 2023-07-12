package main

import (
	"github.com/2023-DSGW-Novel-Engineering/cation-backend/controller"
	"github.com/2023-DSGW-Novel-Engineering/cation-backend/initializers"
	"github.com/2023-DSGW-Novel-Engineering/cation-backend/middleware"
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
	engine.POST("/auth/register", controller.Register)
	engine.POST("/auth/login", controller.Login)
	engine.POST("/auth/logout", middleware.RequireAuth, controller.Logout)
	engine.POST("/auth/vaildate", middleware.RequireAuth, controller.Vaildate)

	engine.GET("/users/:id", controller.GetUserInfo)

	engine.Run(":9190")
}
