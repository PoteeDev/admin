package api

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/PoteeDev/admin/api/database"
	"github.com/PoteeDev/admin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

func GetScenarioFromFile() (*models.Scenario, error) {
	var scenario models.Scenario

	yamlFile, err := ioutil.ReadFile("scenario.yml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &scenario)
	if err != nil {
		return nil, err
	}
	return &scenario, nil
}

func UploadScenarioToMongo(scenario *models.Scenario) {
	col := database.GetCollection(database.DB, "settings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := col.UpdateOne(ctx, bson.M{"id": "scenario"}, bson.M{"$set": scenario}, options.Update().SetUpsert(true))
	if err != nil {
		log.Println("upload to mongo error:", err.Error())
	}
	log.Println("upload result:", result)
}

func UploadScenario() {
	scenario, err := GetScenarioFromFile()
	if err != nil {
		log.Println("load from file error", err.Error())
	}
	log.Println(scenario)
	UploadScenarioToMongo(scenario)
}
