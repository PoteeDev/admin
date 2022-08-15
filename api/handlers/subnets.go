package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PoteeDev/admin/api/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Subnet struct {
	Entity string `json:"entity"`
	Subnet string `json:"subnet"`
}

func AddSubnetToEntity(c *gin.Context) {
	var subnet Subnet
	jsonErr := c.BindJSON(&subnet)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": jsonErr.Error()})
		return
	}

	col := database.GetCollection(database.DB, "entities")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := col.UpdateOne(
		ctx,
		bson.M{"id": subnet.Entity},
		bson.M{"$set": bson.M{"subnet": subnet}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Println("mongo error:", err.Error())
	}
	c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("add %s subnet to user %s", subnet.Subnet, subnet.Entity)})
}
