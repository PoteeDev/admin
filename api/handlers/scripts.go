package handlers

import (
	"context"
	"fmt"
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

const bucketName = "scripts"

func GetScriptsList(c *gin.Context) {
	minioClient := api.ConnectMinio()
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
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
	c.JSON(http.StatusOK, gin.H{bucketName: scripts})

}

func GetScript(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "argument empty: 'name'"})
		return
	}
	minioClient := api.ConnectMinio()
	object, err := minioClient.GetObject(context.Background(), bucketName, name, minio.GetObjectOptions{})
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

func UploadScript(c *gin.Context) {
	script, _ := c.FormFile("file")
	log.Println(script.Filename)

	// Upload the file to specific dst.
	// c.SaveUploadedFile(file, dst)
	file, _ := script.Open()

	minioClient := api.ConnectMinio()
	uploadInfo, err := minioClient.PutObject(context.Background(), bucketName, script.Filename, io.Reader(file), script.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded bytes: ", uploadInfo)

	c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("'%s' uploaded!", script.Filename), "status": "ok"})
}

func DeleteScript(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "argument empty: 'name'"})
		return
	}
	minioClient := api.ConnectMinio()
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	err := minioClient.RemoveObject(context.Background(), bucketName, name, opts)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("'%s' deleted!", name), "status": "ok"})
}
