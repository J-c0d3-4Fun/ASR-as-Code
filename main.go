package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	NewAuthenticator()
	ListBuckets(context.Background())
	buckets, err := ListBuckets(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found", len(buckets), "buckets:")
	for _, bucket := range buckets {
		fmt.Println("-", *bucket.Name)
	}
}
