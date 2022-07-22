package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeamInfo struct {
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	Name      string    `bson:"name"`
	Login     string    `bson:"login"`
	Address   string    `bson:"address"`
	SshPubKey string    `bson:"shh_pub_key"`
}

type DBTeam struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Name      string             `bson:"name"`
	Login     string             `bson:"login"`
	Hash      string             `bson:"hash"`
	Address   string             `bson:"address"`
	SshPubKey string             `bson:"shh_pub_key"`
}

func (t *DBTeam) AddTeam() error {
	col := GetCollection(DB, "teams")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	t.ID = primitive.NewObjectID()
	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return err
	}

	return nil
}

func GetAllTeams() ([]TeamInfo, error) {
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
