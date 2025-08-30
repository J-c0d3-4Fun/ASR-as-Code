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
		policy, err := Iam.ListPolicies(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		//
		fmt.Printf("Found %d IAM policies:\n", len(policy))
		for _, p := range policy {
			fmt.Printf("- %s\n", *p.PolicyName)
		}
		user, err := Iam.ListUsers(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Found %d Users:\n", len(user))
		for _, u := range user {
			fmt.Printf("- %s\n", *u.UserName)
		}

		for _, users := range user {
			mfa, err := Iam.CheckMFA(context.Background(), users.UserName)
			if err != nil {
				fmt.Printf("Error listing MFA devices for user %s: %v\n", *users.UserName, err)
				continue
			}
			if len(mfa) > 0 {
				fmt.Printf("User %s: MFA Enabled\n", *users.UserName)
			} else {
				fmt.Printf("User %s: MFA Not Enabled\n", *users.UserName)

			}

		}

	},
}

func init() {

	rootCmd.AddCommand(getIamPolicy)

}
