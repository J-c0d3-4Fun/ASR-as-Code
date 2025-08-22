package cmd

import (
	Iam "ASR-as-Code/pkg/IAM"
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var getIamPolicy = &cobra.Command{
	Use:     "get-iam",
	Aliases: []string{"iam"},
	Short:   "List and analyze iam policies for security misconfigurations",
	Long:    "",

	Run: func(cmd *cobra.Command, args []string) {
		Iam.ListPolicies(context.Background())
		policy, err := Iam.ListPolicies(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		//
		fmt.Printf("Found %d IAM policies:\n", len(policy))
		for _, p := range policy {
			fmt.Printf("- %s\n", *p.PolicyName)
		}

	},
}

func init() {

	rootCmd.AddCommand(getIamPolicy)

}
