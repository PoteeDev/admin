package handlers

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/PoteeDev/admin/api"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

type Script struct {
	Name    string `json:"name"`
	Size    int64  `json:"size"`
	Modifed string `json:"modifed"`
}

type ScriptName struct {
	Name string `json:"name"`
}

func GetScriptsList(c *gin.Context) {
	minioClient := api.ConnectMinio()
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	objectCh := minioClient.ListObjects(ctx, "scripts", minio.ListObjectsOptions{
		Prefix:    "",
		Recursive: true,
	})
	var scripts []Script
	for object := range objectCh {
		if object.Err != nil {
			log.Println(object.Err)
			continue
		}
		scripts = append(scripts, Script{
			Name:    object.Key,
			Size:    object.Size,
			Modifed: object.LastModified.String(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"scripts": scripts})

}

func GetScript(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "argument empty: 'name'"})
		return
	}
	minioClient := api.ConnectMinio()
	object, err := minioClient.GetObject(context.Background(), "scripts", name, minio.GetObjectOptions{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	defer object.Close()
	file, err := io.ReadAll(object)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	c.Data(http.StatusOK, "text/plain", file)
}

func UploadScript() {

}

func DeleteScript() {

}
