package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type Auth struct {
	Cfg aws.Config
	Iam *iam.Client
}

func NewAuthenticator() (*Auth, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err

	}
	return &Auth{
		Cfg: cfg,
		Iam: iam.NewFromConfig(cfg),
	}, nil

}
