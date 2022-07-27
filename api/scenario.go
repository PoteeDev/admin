package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
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

func ConvertTime(time string) (int, error) {
	switch string(time[len(time)-1]) {
	case "s":
		seconds, err := strconv.Atoi(time[:len(time)-1])
		return seconds, err
	case "m":
		seconds, err := strconv.Atoi(time[:len(time)-1])
		return seconds * 60, err
	case "h":
		seconds, err := strconv.Atoi(time[:len(time)-1])
		return seconds * 60 * 60, err
	default:
		return 0, fmt.Errorf("invalid time suffix: not 's' or 'm'")
	}
}

func GenerateRounds(scenario *models.Scenario) {
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)
	roundPeriod, err := ConvertTime(scenario.Period)
	if err != nil {
		log.Fatal(err)
	}
	totalTime, err := ConvertTime(scenario.Time)
	if err != nil {
		log.Fatal(err)
	}
	for serviceName, service := range scenario.Services {
		for exploitName, exploit := range service.Exploits {
			exploitPeriod, err := ConvertTime(exploit.Period)
			if err != nil {
				log.Fatal(err)
			}
			exploitRoundsInterval := exploitPeriod / roundPeriod
			var rounds = []int{exploit.Round}
			for round := exploit.Round + exploitRoundsInterval; round < totalTime/roundPeriod; round += exploitRoundsInterval {
				rounds = append(rounds, r.Intn(exploitRoundsInterval)+round)
			}
			exploit.Rounds = rounds
			log.Printf("generate rounds for %s:%s %v", serviceName, exploitName, rounds)
			scenario.Services[serviceName].Exploits[exploitName] = exploit

		}
	}
}

func UploadScenario() {
	scenario, err := GetScenarioFromFile()
	if err != nil {
		log.Println("load from file error", err.Error())
	}
	GenerateRounds(scenario)
	log.Println(scenario)
	UploadScenarioToMongo(scenario)
}
