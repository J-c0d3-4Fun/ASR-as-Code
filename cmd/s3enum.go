package cmd

import (
	"ASR-as-Code/pkg/s3"

	"github.com/spf13/cobra"
)

var getBucketsCmd = &cobra.Command{
	Use:     "get-buckets",
	Aliases: []string{"s3"},
	Short:   "List and analyze S3 buckets for security misconfigurations",
	Long:    "",

	Run: func(cmd *cobra.Command, args []string) {
		s3.Auth.

	},
}

func init() {
	rootCmd.AddCommand(getBucketsCmd)
}
