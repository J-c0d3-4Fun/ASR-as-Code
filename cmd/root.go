package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "asr",
	Short: "asr (Attack Surface Reduction as Code) is a Lightweight tool to scan AWS resources for security misconfiguration",
	Long: `The asr (Attack Surface Reduction as Code) CLI helps identify and reduce
cloud misconfigurations directly from your terminal.  

Currently supports AWS S3 checks:
  • Detects buckets without encryption
  • Verifies block public access settings
  • Analyzes attached bucket policies

Future releases will expand coverage to other AWS services.

Use asr to bring security into your workflow early by running checks
locally, in CI/CD pipelines, or as part of automated audits.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// Handles the execution of the root command and handles possible errors by printing the error messages to the console
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
