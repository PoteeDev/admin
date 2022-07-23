package main

import (
	"github.com/PoteeDev/admin/api/database"
	"github.com/PoteeDev/admin/api/handlers"
	"github.com/PoteeDev/auth/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDB()
	database.CreateAdmin()
	r.GET("/ping", middleware.AdminAuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/teams", middleware.AdminAuthMiddleware(), handlers.TeamsList)
	r.Run()
}
