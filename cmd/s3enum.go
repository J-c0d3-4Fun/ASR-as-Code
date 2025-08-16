package cmd

import (
	"ASR-as-Code/pkg/s3"
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var getBucketsCmd = &cobra.Command{
	Use:     "get-buckets",
	Aliases: []string{"s3"},
	Short:   "List and analyze S3 buckets for security misconfigurations",
	Long:    "",

	Run: func(cmd *cobra.Command, args []string) {

		s3.ListBuckets(context.Background())
		buckets, err := s3.ListBuckets(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Found", len(buckets), "buckets:")
		for _, bucket := range buckets {
			fmt.Println("-", *bucket.Name)
		}
		findings, err := s3.BucketResults()
		if err != nil {
			log.Fatal(err)
		}
		for _, finding := range findings {
			fmt.Printf("\nBucket: %s\n Public: %v\n Encrypted: %v\n BucketPolicy: %v\n",
				finding.Name, finding.IsPublic, finding.HasEncryption, finding.BucketPolicy)
		}

	},
}

func init() {

	rootCmd.AddCommand(getBucketsCmd)

}
