package handlers

import (
	"net/http"

	"github.com/PoteeDev/admin/api/database"
	"github.com/PoteeDev/entities/vpn"
	"github.com/gin-gonic/gin"
)

func GenerateVpns(c *gin.Context) {
	enities, err := database.GetAllEntities()
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	for _, entity := range enities {
		if entity.Visible {
			client := vpn.CreateVpnCLient(entity.Login)
			client.CreateConfig()
			client.AddRoute(entity.Subnet)
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "msg": "vpn created"})
}
