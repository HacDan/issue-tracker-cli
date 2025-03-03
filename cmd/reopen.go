package cmd

import (
	"fmt"

	"github.com/hacdan/issue-tracker-cli/storage"
	"github.com/hacdan/issue-tracker-cli/utils"
	"github.com/spf13/cobra"
)

// reopenCmd represents the reopen command
var reopenCmd = &cobra.Command{
	Use:   "reopen",
	Short: "reopen opens a closed issue by Id number",
	Long: `reopen opens any closed issued that was previously closed
The issue is identified by ide number. For example

issue reopen 42
> Successfully reopened issue: 42`,
	Run: func(cmd *cobra.Command, args []string) {
		issueId := utils.StringToInt(args[0])
		storage := storage.NewStorage()

		storage.ReopenIssue(issueId)
		fmt.Println("Successfully reopened issue: ", issueId)
	},
}

func init() {
	rootCmd.AddCommand(reopenCmd)
}
