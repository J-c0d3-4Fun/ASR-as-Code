package iam

import (
	"context"
	"fmt"
	"log"

	auth "ASR-as-Code/pkg/auth"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

// ListPolicies gets up to maxPolicies policies.
func ListPolicies(ctx context.Context) ([]types.Policy, error) {
	var policies []types.Policy
	result, err := Sess.IAM.ListPolicies(ctx, &iam.ListPoliciesInput{})
	if err != nil {
		log.Printf("Couldn't list policies. Here's why: %v\n", err)
	} else {
		policies = result.Policies
	}
	return policies, err
}

// List all the current users
func ListUsers(ctx context.Context) ([]types.User, error) {
	var users []types.User
	listUsers, err := Sess.IAM.ListUsers(context.TODO(), &iam.ListUsersInput{})
	if err != nil {
		fmt.Printf("Error listing users: %v\n", err)
	} else {
		users = listUsers.Users
	}
	return users, err

}

func CheckMFA(ctx context.Context, userNames *string) ([]types.MFADevice, error) {
	var mfa []types.MFADevice
	getMfa, err := Sess.IAM.ListMFADevices(context.TODO(), &iam.ListMFADevicesInput{
		UserName: userNames,
	})
	if err != nil {
		return nil, err
	} else {
		mfa = getMfa.MFADevices
	}
	return mfa, nil

}

// TODO create a function that sets the context the Background

var Sess *auth.Auth

// initializes the auth.go first allows us to use the AWS creds
func init() {
	var err error
	Sess, err = auth.NewAuthenticator()
	if err != nil {
		log.Fatal(err)
	}

}
