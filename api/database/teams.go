package database

import (
	"context"
	"time"

	"github.com/PoteeDev/team/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTeams() ([]models.TeamInfo, error) {
	var teams []models.TeamInfo
	col := GetCollection(DB, "teams")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	results, err := col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer results.Close(ctx)
	defer results.Close(ctx)
	for results.Next(ctx) {
		var team models.TeamInfo
		if err = results.Decode(&team); err != nil {
			return nil, err
		}

		teams = append(teams, team)
	}
	return teams, err
}
