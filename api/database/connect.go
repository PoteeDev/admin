package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/PoteeDev/team/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func ConnectDB() *mongo.Client {
	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s", "admin", os.Getenv("MONGO_PASS"), os.Getenv("MONGO_HOST"))
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
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
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
func CreateAdmin() {
	col := GetCollection(DB, "teams")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	hash, err := hashPassword(os.Getenv("ADMIN_PASS"))
	if err != nil {
		log.Fatalln("Create admin error", err.Error())
	}
	admin := models.Team{
		Name:      "Admin",
		Login:     "admin",
		Hash:      hash,
		UpdatedAt: time.Now(),
		Visible:   false,
		Blocked:   false,
	}
	result, err := col.UpdateOne(ctx,
		bson.M{"login": "admin"},
		bson.M{"$set": admin},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Fatalln("Create admin error", err.Error())
	}
	log.Println("Admin created", result.MatchedCount)
}
