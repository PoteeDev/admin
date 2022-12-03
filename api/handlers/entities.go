package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PoteeDev/admin/api/database"
	entityDatabase "github.com/PoteeDev/entities/database"
	"github.com/PoteeDev/entities/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/gosimple/slug"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func EntitiesList(c *gin.Context) {
	entities, err := database.GetAllEntities()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"teams": entities})
}

type Entity struct {
	Name     string `json:"name" validate:"required,min=3,max=30"`
	Password string `json:"password" validate:"required,min=8,max=100"`
	Subnet   string `json:"subnet"`
	IP       string `json:"ip"`
	Visible  bool   `json:"visible"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (t *Entity) WriteEntity(login string) error {

	// ipAddress := generateIp(len(teams) + 1)
	hash, hashErr := HashPassword(t.Password)
	if hashErr != nil {
		return hashErr
	}
	col := database.GetCollection(database.DB, "entities")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := col.InsertOne(ctx, &models.Entity{
		Name:      t.Name,
		Login:     login,
		Hash:      hash,
		Visible:   t.Visible,
		Subnet:    t.Subnet,
		IP:        t.IP,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}

func CreateEntity(c *gin.Context) {
	var entity Entity
	jsonErr := c.BindJSON(&entity)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": jsonErr.Error()})
		return
	}
	validate := validator.New()
	if err := validate.Struct(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}
	// create slug name for entity
	login := slug.Make(entity.Name)
	// todo: change name policy in vpn service and remove this line
	login = strings.Replace(login, "-", "_", -1)

	// check if user already exists
	// todo: create function in mongo to check exists usern, not find
	dbEntity, err := entityDatabase.GetEntity(login)
	fmt.Println(dbEntity, err)
	if dbEntity != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "entity already exists"})
		return
	}
	// write user to database
	if writeErr := entity.WriteEntity(login); writeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": writeErr.Error()})
	}
	// if all ok - return message and entity login
	c.JSON(http.StatusOK, gin.H{
		"msg":   fmt.Sprintf("The entity %s created", entity.Name),
		"login": login,
	})
}

func DeleteEntities(c *gin.Context) {
	name, ok := c.Params.Get("name")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "argument empty: 'name'"})
		return
	}
	col := database.GetCollection(database.DB, "entities")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := col.DeleteOne(ctx, bson.M{"name": name})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("%s deleted!", name)})
}
