package handlers

import (
	"net/http"

	"github.com/PoteeDev/admin/api/database"
	"github.com/gin-gonic/gin"
)

func EntitiesList(c *gin.Context) {
	entities, err := database.GetAllEntities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"teams": entities})
}
