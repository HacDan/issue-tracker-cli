package cmd

import (
	"fmt"

	"github.com/hacdan/issue-tracker-cli/storage"
	"github.com/hacdan/issue-tracker-cli/types"
	"github.com/hacdan/issue-tracker-cli/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a new issue to the tracker",
	Long: `The add command adds a new issue to the traacker. The title, description, and user are required fields.
the rest will be filled with sane defaults. Examples:

itc add -t "Example Title" -d "Example Description" -p "low" -s "inprogress" -a "John"
> Issue added successfully
itc add -t "Example Title" -d Example Description" -a "Jim"
> Issue added successfully`,
	Run: func(cmd *cobra.Command, args []string) {
		issue := types.Issue{}
		storage := storage.NewStorage()

		issue.Title, _ = cmd.Flags().GetString("title")
		issue.Description, _ = cmd.Flags().GetString("description")
		sPriority, _ := cmd.Flags().GetString("priority")
		issue.Priority = utils.StringToPriority(sPriority)
		sStatus, _ := cmd.Flags().GetString("status")
		issue.Status = utils.SStatusToStatus(sStatus)
		issue.User, _ = cmd.Flags().GetString("user")

		_, err := storage.AddIssue(issue)
		if err != nil {
			fmt.Println("Error adding issue: ", err)
		}
		fmt.Println("Issue added successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
