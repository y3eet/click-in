package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/y3eet/click-in/internal/config"
)

func NewMinioClient() (*s3.Client, error) {
	cfg := config.Cfg
	minioCfg, err := awsconfig.LoadDefaultConfig(
		context.TODO(),
		awsconfig.WithRegion("us-east-1"),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.MinioRootUser,
				cfg.MinioRootPassword,
				"",
			),
		),
	)
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(minioCfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(cfg.MinioEndpoint)
		o.UsePathStyle = true // VERY important for MinIO
	}), nil
}
