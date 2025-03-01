/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "issue-tracker-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.issue-tracker-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	var title string
	rootCmd.PersistentFlags().StringVarP(&title, "title", "t", "", "Title for issue")

	var description string
	rootCmd.PersistentFlags().StringVarP(&description, "description", "d", "", "Description for issue")

	var priority string
	rootCmd.PersistentFlags().StringVarP(&priority, "priority", "p", "low", "Priority for issue")

	var status string
	rootCmd.PersistentFlags().StringVarP(&status, "status", "s", "open", "Status for issue")

	var user string
	rootCmd.PersistentFlags().StringVarP(&user, "user", "a", "", "User assigned to issue")
}
