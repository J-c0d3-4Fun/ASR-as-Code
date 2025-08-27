package auth

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Auth struct {
	Cfg aws.Config
	S3  *s3.Client
	IAM *iam.Client
}

func NewAuthenticator() (*Auth, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err

	}
	return &Auth{
		Cfg: cfg,
		S3:  s3.NewFromConfig(cfg),
		IAM: iam.NewFromConfig(cfg),
	}, nil

}
