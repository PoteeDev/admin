package handlers

import (
	"net/http"

	"github.com/PoteeDev/admin/api/database"
	"github.com/gin-gonic/gin"
)

func TeamsList(c *gin.Context) {
	teams, err := database.GetAllTeams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"teams": teams})
}
