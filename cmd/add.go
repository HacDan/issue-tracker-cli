/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/hacdan/issue-tracker-cli/storage"
	"github.com/hacdan/issue-tracker-cli/types"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		issue := types.Issue{}
		issue.Title, _ = cmd.Flags().GetString("title")
		issue.Description, _ = cmd.Flags().GetString("description")
		sStatus, _ := cmd.Flags().GetString("status")
		issue.Status = setStatus(sStatus)
		issue.User, _ = cmd.Flags().GetString("user")

		storage := storage.NewStorage()
		_, err := storage.AddIssue(issue)
		if err != nil {
			fmt.Println("Error adding issue: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	var title string
	addCmd.Flags().StringVarP(&title, "title", "t", "", "Title for issue")

	var description string
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description for issue")

	var priority string
	addCmd.Flags().StringVarP(&priority, "priority", "p", "low", "Priority for issue")

	var status string
	addCmd.Flags().StringVarP(&status, "status", "s", "open", "Status for issue")

	var user string
	addCmd.Flags().StringVarP(&user, "user", "a", "", "User assigned to issue")
}

func setStatus(s string) types.IssueStatus {
	s = strings.ToLower(s)
	if s == "open" {
		return types.OPEN
	}
	if s == "closed" {
		return types.CLOSED
	}
	if s == "inprogress" {
		return types.INPROGRESS
	}
	return types.OPEN
}
