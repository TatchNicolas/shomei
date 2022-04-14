/*
Copyright © 2022 TatchNicolas

*/
package cmd

import (
	"os"

	"github.com/TatchNicolas/shomei/cmd/aws"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shomei",
	Short: "Prints HTTP request header for public cloud API",
	Long: `shomei is a CLI to print HTTP request header for public cloud API to be used with curl or httpie.

This tool assumes credentials for the official CLI is already available in the terminal session.
For example, AWS_XXX environment variables for AWS CLI or ADC for gcloud command.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.shomei.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(aws.Cmd)
}
