package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri).SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

//Client instance
var DB *mongo.Client = ConnectDB()

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("ad").Collection(collectionName)
	return collection
}
