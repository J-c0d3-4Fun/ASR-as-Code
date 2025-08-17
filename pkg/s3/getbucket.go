package s3

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
)

type BucketFindings struct {
	Name               string
	IsPublic           bool
	HasEncryption      bool
	BucketPolicyStatus string
	BucketPolicy       string
}

func ListBuckets(ctx context.Context) ([]types.Bucket, error) {
	// calls Auth struct in auth.go
	auth, err := NewAuthenticator()
	if err != nil {
		return nil, err
	}
	var output *s3.ListBucketsOutput
	var buckets []types.Bucket
	bucketPaginator := s3.NewListBucketsPaginator(auth.S3, &s3.ListBucketsInput{})
	for bucketPaginator.HasMorePages() {
		output, err = bucketPaginator.NextPage(ctx)
		if err != nil {
			var apiErr smithy.APIError
			if errors.As(err, &apiErr) && apiErr.ErrorCode() == "AccessDenied" {
				fmt.Println("You don't have permission to list buckets for this account.")
				err = apiErr
			} else {
				log.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
			}
			break
		} else {
			buckets = append(buckets, output.Buckets...)
		}
	}
	return buckets, err
}

func GetBucketPublicAccess(bucketName string) (bool, error) {
	// Call GetPublicAccessBlock for this specific bucket
	// Return true if public, false if blocked

	if bucketName == "" {
		return false, errors.New("bucket name cannot be empty")
	}
	output, err := auth.S3.GetPublicAccessBlock(context.TODO(), &s3.GetPublicAccessBlockInput{
		Bucket: &bucketName,
	})
	if err != nil {
		return false, err
	}
	if output.PublicAccessBlockConfiguration != nil {
		return false, nil
	}
	return true, nil
}

func GetBucketEncryption(bucketName string) (bool, error) {
	// Checks bucketname to see if it empty
	if bucketName == "" {
		return false, errors.New("bucket name cannot be empty")
	}
	output, err := auth.S3.GetBucketEncryption(context.TODO(), &s3.GetBucketEncryptionInput{
		Bucket: &bucketName,
	})
	if err != nil {
		return false, err
	}
	if output.ServerSideEncryptionConfiguration != nil {
		return true, nil
	}
	return false, nil

}
func GetBucketPolicyStatus(bucketName string) (string, error) {
	if bucketName == "" {
		return "", errors.New("bucket name cannot be empty")
	}
	output, err := auth.S3.GetBucketPolicyStatus(context.TODO(), &s3.GetBucketPolicyStatusInput{
		// bucket parameter for call
		Bucket: &bucketName,
	})

	if err != nil {
		return "", nil
	}
	if output.PolicyStatus != nil {
		// converting to string
		return fmt.Sprint(output.PolicyStatus.IsPublic), nil
	}
	return "", nil

}

func GetBucketPolicy(bucketName string) (string, error) {
	if bucketName == "" {
		return "", errors.New("bucket name cannot be empty")
	}
	output2, err := auth.S3.GetBucketPolicy(context.TODO(), &s3.GetBucketPolicyInput{
		// Must set the bucket name to make a call. converting the cunftion parmaeter to the required pointer type
		Bucket: &bucketName,
	})

	if err != nil {
		return "", err
	}
	if output2.Policy != nil {
		return *output2.Policy, nil

	}
	return "No Policy", nil
}

func GetBucketRegion(bucketName string) (string, error) {
	output, err := auth.S3.GetBucketLocation(context.TODO(), &s3.GetBucketLocationInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatalf("failed to get bucket location, %v", err)
	}
	return string(output.LocationConstraint), nil

}

func BucketResults() ([]BucketFindings, error) {
	// create a array of the object BucketFidings
	var findings []BucketFindings
	results, err := ListBuckets(context.Background())
	if err != nil {
		return nil, err
	}
	// Looping through the object/function ListBuckets of the variable results
	for _, value := range results {
		// call the function bucketRegion to get the current credentials region
		bucketRegion, err := GetBucketRegion(*value.Name)
		if err != nil {
			continue
		}
		// skips the buckets that dont match the current region of the user
		if bucketRegion != auth.Cfg.Region {
			continue
		}
		b := BucketFindings{}

		// creating the instances of the BucketFindings
		b.Name = *value.Name
		b.HasEncryption, err = GetBucketEncryption(*value.Name)
		if err != nil {
			return nil, err
		}
		b.IsPublic, err = GetBucketPublicAccess(*value.Name)
		if err != nil {
			return nil, err
		}
		findings = append(findings, b)

		b.BucketPolicyStatus, err = GetBucketPolicyStatus(*value.Name)
		if err != nil {
			return nil, err
		}
		for _, value := range results {
			// call the function bucketRegion to get the current credentials region
			bucketRegion, err := GetBucketRegion(*value.Name)
			if err != nil {
				continue
			}
			// skips the buckets that dont match the current region of the user
			if bucketRegion != auth.Cfg.Region {
				continue
			}

			b.BucketPolicy, err = GetBucketPolicy(*value.Name)

			if err != nil {
				// Log error but continue to next bucket
				fmt.Printf("Error getting policy for %s: %v\n", *value.Name, err)
				b.BucketPolicy = "Error"
				// Don't return - keep processing other buckets
			}
		}

	}
	return findings, nil
}
