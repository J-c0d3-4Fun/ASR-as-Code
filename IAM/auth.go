package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type Auth struct {
	Cfg aws.Config
}

func NewAuthenticator() (*Auth, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err

	}
	return &Auth{
		Cfg: cfg,
	}, nil

}
