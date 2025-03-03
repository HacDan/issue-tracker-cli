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
	Long: `The list command lists all issues, with filtering options
	You can filter on status, priority, or user. For example:

issue list 
issue list -s open               # Show only open issues  
issue list -p high               # Show high-priority issues  
issue list -a alice              # Show issues assigned to Alice`,
	Run: func(cmd *cobra.Command, args []string) {
		status, _ := cmd.Flags().GetString("status")
		user, _ := cmd.Flags().GetString("user")
		priority, _ := cmd.Flags().GetString("priority")

		if cmd.Flags().Changed("status") {
			store := storage.NewStorage()
			issues, _ := store.GetIssueByStatus(status)
			for _, issue := range issues {
				issue.Print()
			}
			return
		} else if cmd.Flags().Changed("user") {
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
}
