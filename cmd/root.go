/*
Copyright © 2022 TatchNicolas

*/
package cmd

import (
	"os"

	"github.com/TatchNicolas/shomei/cmd/aws"
	"github.com/TatchNicolas/shomei/cmd/playground"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shomei",
	Short: "Prints HTTP request header for public cloud API",
	Long: `shomei is a CLI to print HTTP request header for public cloud API to be used with curl or httpie.

This tool assumes credentials for the official CLI is already available in the terminal session.
For example, AWS_XXX environment variables for AWS CLI or ADC for gcloud command.
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(aws.Cmd)
	rootCmd.AddCommand(playground.Cmd)
}
