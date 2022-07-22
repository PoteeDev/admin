package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type TeamInfo struct {
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	Name      string    `bson:"name"`
	Login     string    `bson:"login"`
	Address   string    `bson:"address"`
	SshPubKey string    `bson:"shh_pub_key"`
}

func List() ([]TeamInfo, error) {
	var teams []TeamInfo
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
		var team TeamInfo
		if err = results.Decode(&team); err != nil {
			return nil, err
		}

		teams = append(teams, team)
	}
	return teams, err
}
