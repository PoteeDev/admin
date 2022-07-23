package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/PoteeDev/admin/api"
	"github.com/PoteeDev/admin/api/database"
	"github.com/PoteeDev/admin/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func GetScenario(c *gin.Context) {
	var scenario models.Scenario
	col := database.GetCollection(database.DB, "settings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := col.FindOne(ctx, bson.M{"id": "scenario"}).Decode(&scenario)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"scenario": ""})
	}
	c.JSON(http.StatusOK, gin.H{"scenario": scenario})
}

func UpdateScenario(c *gin.Context) {
	var s models.JsonScenario
	if err := c.BindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	api.UploadScenarioToMongo(&s.Scenario)
	GetScenario(c)
}
