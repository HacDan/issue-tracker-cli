/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hacdan/issue-tracker-cli/storage"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all open issues for a project",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		status, _ := cmd.Flags().GetString("status")
		user, _ := cmd.Flags().GetString("user")
		priority, _ := cmd.Flags().GetString("priority")

		if cmd.Flags().Changed("status") {
			fmt.Println("Status called: ", status)
			store := storage.NewStorage()
			issues, _ := store.GetIssueByStatus(status)
			for _, issue := range issues {
				issue.Print()
			}
			return
		} else if cmd.Flags().Changed("user") {
			fmt.Println("User called: ", user)
			store := storage.NewStorage()
			issues, err := store.GetIssueByUser(user)
			if err != nil {
				fmt.Println(err)
			}
			for _, issue := range issues {
				issue.Print()
			}
			return
		} else if cmd.Flags().Changed("priority") {
			store := storage.NewStorage()

			issues, err := store.GetIssueByPriority(priority)
			if err != nil {
				fmt.Println(err)
			}

			for _, issue := range issues {
				issue.Print()
			}

			return
		} else {
			store := storage.NewStorage()
			issues, _ := store.GetIssues()
			for _, issue := range issues {
				issue.Print()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
