package handlers

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/y3eet/click-in/internal/storage"
)

const (
	MaxFileSize = 10 * 1024 * 1024 // 10 MB
	BucketName  = "bucket"
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
	if file.Size > MaxFileSize {
		msg := fmt.Sprintf("Max file size is %d MB", MaxFileSize/(1024*1024))
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	defer src.Close()

	client, err := storage.NewMinioClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	key := filepath.Base(file.Filename + "-" + time.Now().Format("20060102150405"))

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(BucketName),
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

func ViewFile(c *gin.Context) {
	client, err := storage.NewMinioClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	key := c.Param("key")

	resp, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	c.DataFromReader(http.StatusOK, *resp.ContentLength, *resp.ContentType, resp.Body, nil)
}
