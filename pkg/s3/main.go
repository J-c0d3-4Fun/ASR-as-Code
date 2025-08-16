package s3

import (
	"context"
	"fmt"
	"log"
)

// Package variable
var auth *Auth

func main() {
	var err error
	auth, err = NewAuthenticator()
	if err != nil {
		log.Fatal(err)
	}

	buckets, err := ListBuckets(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found", len(buckets), "buckets:")
	for _, bucket := range buckets {
		fmt.Println("-", *bucket.Name)
	}
	findings, err := BucketResults()
	if err != nil {
		log.Fatal(err)
	}
	for _, finding := range findings {
		fmt.Printf("\nBucket: %s\n Public: %v\n Encrypted: %v\n BucketPolicy: %v\n",
			finding.name, finding.isPublic, finding.hasEncryption, finding.bucketPolicy)
	}

}
