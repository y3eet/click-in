package handlers

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/y3eet/click-in/internal/storage"
)

func FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	client, err := storage.NewMinioClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	key := filepath.Base(file.Filename)

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String("bucket"),
		Key:         aws.String(key),
		Body:        src,
		ContentType: aws.String(file.Header.Get("Content-Type")),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload successful",
		"file":    key,
	})
}
