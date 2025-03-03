package cmd

import (
	"fmt"

	"github.com/hacdan/issue-tracker-cli/storage"
	"github.com/spf13/cobra"
)

// closeCmd represents the close command
var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "The close command closes the specified issue by ID",
	Long: `Used to close issues by Id number. 
For example:

issue close 42`,
	Run: func(cmd *cobra.Command, args []string) {
		issueId := stringToInt(args[0])
		storage := storage.NewStorage()

		err := storage.CloseIssue(issueId)
		if err != nil {
			fmt.Println("Error updating issue: ", err)
		}
		fmt.Println("Successfully closed issue number: ", issueId)
	},
}

func init() {
	rootCmd.AddCommand(closeCmd)
}
