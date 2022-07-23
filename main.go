package main

import (
	"github.com/PoteeDev/admin/api"
	"github.com/PoteeDev/admin/api/database"
	"github.com/PoteeDev/admin/api/handlers"
	"github.com/PoteeDev/auth/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDB()
	api.CreateAdmin()
	api.UploadScenario()
	api.UploadScripts()
	r.GET("/ping", middleware.AdminAuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/teams", middleware.AdminAuthMiddleware(), handlers.TeamsList)
	r.GET("/scenario", middleware.AdminAuthMiddleware(), handlers.GetScenario)
	r.POST("/scenario/update", middleware.AdminAuthMiddleware(), handlers.UpdateScenario)
	r.GET("/scripts", middleware.AdminAuthMiddleware(), handlers.GetScriptsList)
	r.GET("/scripts/get", middleware.AdminAuthMiddleware(), handlers.GetScript)
	r.Run()
}
