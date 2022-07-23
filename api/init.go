package api

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/PoteeDev/admin/api/database"
	"github.com/PoteeDev/team/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CreateAdmin() {
	col := database.GetCollection(database.DB, "teams")
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
