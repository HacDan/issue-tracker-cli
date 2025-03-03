package cmd

import (
	"strings"

	"github.com/hacdan/issue-tracker-cli/storage"
	"github.com/hacdan/issue-tracker-cli/types"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edits the specified issue by Id number",
	Long: `edit is used to modify the specified issue
identified by Id. This command can be used to edit any field 
in the issue. Example: 

issue edit 42 -t "Fix login error" -p critical -s in-progress`,
	Run: func(cmd *cobra.Command, args []string) {
		store := storage.NewStorage()
		updatedIssue := types.Issue{}

		updatedIssue.Id = stringToInt(args[0])

		if cmd.Flags().Changed("title") {
			updatedIssue.Title, _ = cmd.Flags().GetString("title")
		}
		if cmd.Flags().Changed("description") {
			updatedIssue.Description, _ = cmd.Flags().GetString("description")
		}
		if cmd.Flags().Changed("priority") {
			sPriority, _ := cmd.Flags().GetString("priority")
			updatedIssue.Priority = stringToPriority(sPriority)
		}
		if cmd.Flags().Changed("status") {
			updatedStatus, _ := cmd.Flags().GetString("status")
			updatedIssue.Status = sStatusToStatus(updatedStatus)
		}
		if cmd.Flags().Changed("user") {
			updatedIssue.User, _ = cmd.Flags().GetString("user")
		}

		store.UpdateIssue(updatedIssue)

	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}

// TODO: Refactor into utilities package
func sStatusToStatus(s string) types.IssueStatus {
	if strings.ToLower(s) == "open" {
		return types.OPEN
	}
	if strings.ToLower(s) == "inprogress" {
		return types.INPROGRESS
	}
	if strings.ToLower(s) == "closed" {
		return types.CLOSED
	}
	return types.OPEN
}
