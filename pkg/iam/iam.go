package Iam

import (
	"context"
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

var Sess *auth.Auth

// initializes the auth.go first allows us to use the AWS creds
func init() {
	var err error
	Sess, err = auth.NewAuthenticator()
	if err != nil {
		log.Fatal(err)
	}

}
