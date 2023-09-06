package database

import (
	"context"
	"time"

	"github.com/PoteeDev/entities/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllEntities() ([]models.EntityInfo, error) {
	var entities []models.EntityInfo
	col := GetCollection(DB, "entities")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	results, err := col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer results.Close(ctx)
	for results.Next(ctx) {
		var entity models.EntityInfo
		if err = results.Decode(&entity); err != nil {
			return nil, err
		}

		entities = append(entities, entity)
	}
	return entities, err
}

func GetEntities() ([]models.EntityInfo, error) {
	var entities []models.EntityInfo
	col := GetCollection(DB, "entities")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	results, err := col.Find(ctx, bson.M{
		"login": bson.M{"$ne": "admin"},
	})
	if err != nil {
		return nil, err
	}
	defer results.Close(ctx)
	for results.Next(ctx) {
		var entity models.EntityInfo
		if err = results.Decode(&entity); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, err
}
