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
	entities := r.Group("/entities")
	{
		entities.GET("/list", middleware.AdminAuthMiddleware(), handlers.EntitiesList)
		entities.POST("/registration", middleware.AdminAuthMiddleware(), handlers.CreateEntity)
		entities.POST("/subnet", middleware.AdminAuthMiddleware(), handlers.AddSubnetToEntity)
		entities.DELETE("/delete", middleware.AdminAuthMiddleware(), handlers.EntitiesList)
		entities.GET("/generate/vpn", middleware.AdminAuthMiddleware(), handlers.GenerateVpns)
	}

	r.GET("/scenario", middleware.AdminAuthMiddleware(), handlers.GetScenario)
	r.POST("/scenario/update", middleware.AdminAuthMiddleware(), handlers.UpdateScenario)

	scripts := r.Group("/scripts", middleware.AdminAuthMiddleware())
	{
		scripts.GET("/list", handlers.GetScriptsList)
		scripts.GET("/get", handlers.GetScript)
		scripts.POST("/upload", handlers.UploadScript)
		scripts.DELETE("/delete", handlers.DeleteScript)
	}
	r.Run()
}
