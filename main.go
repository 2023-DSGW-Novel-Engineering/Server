package main

import (
	"github.com/2023-DSGW-Novel-Engineering/cation-backend/controller"
	"github.com/2023-DSGW-Novel-Engineering/cation-backend/initializers"
	"github.com/2023-DSGW-Novel-Engineering/cation-backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	engine := gin.Default()
	engine.Use(middleware.CORSMiddleware())
	engine.Use(cors.Default())

	// userController
	engine.POST("/auth/register", controller.Register)
	engine.POST("/auth/login", controller.Login)
	engine.POST("/auth/logout", middleware.RequireAuth, controller.Logout)
	engine.POST("/auth/vaildate", middleware.RequireAuth, controller.Vaildate)

	engine.GET("/api/users/info", middleware.RequireAuth, controller.GetUserInfo)
	// engine.POST("/api/addfriend", middleware.RequireAuth, controller.AddFriend)

	engine.Run(":4000")
}
