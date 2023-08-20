package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	credential := options.Credential{
		AuthSource: os.Getenv("MONGO_DB"),
		Username:   os.Getenv("MONGO_USER"),
		Password:   os.Getenv("MONGO_PASS"),
	}
	mongoUri := fmt.Sprintf("mongodb://%s", os.Getenv("MONGO_HOST"))
	clientOpts := options.Client().ApplyURI(mongoUri).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(os.Getenv("MONGO_DB")).Collection(collectionName)
	return collection
}
