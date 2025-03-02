/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hacdan/issue-tracker-cli/storage"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")

		fmt.Println(title, description)
		if len(args) > 0 {
			storage := storage.NewStorage()
			issues, err := storage.SearchIssuesByText(args[0])
			if err != nil {
				fmt.Println(err)
			}
			for _, issue := range issues {
				issue.Print()
			}
		}
		if cmd.Flags().Changed("title") {
			storage := storage.NewStorage()
			issues, err := storage.SearchIssuesByTitle(title)
			if err != nil {
				fmt.Println(err)
			}

			for _, issue := range issues {
				issue.Print()
			}
			return
		}
		if cmd.Flags().Changed("description") {
			storage := storage.NewStorage()
			issues, err := storage.SearchIssuesByDescription(description)
			if err != nil {
				fmt.Println(err)
			}
			for _, issue := range issues {
				issue.Print()
			}
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
