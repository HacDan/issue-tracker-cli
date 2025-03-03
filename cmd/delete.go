package cmd

import (
	"fmt"
	"strconv"

	"github.com/hacdan/issue-tracker-cli/storage"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a specified issue by Id number",
	Long: `Deletes a specified issue by Id number. Can also be used for
deleting all-closed issues. Examples:

issue delete 42
issue delete --all-closed       # Delete all closed issues  `,
	Run: func(cmd *cobra.Command, args []string) {
		storage := storage.NewStorage()
		deleteAllClosed, _ := cmd.Flags().GetBool("all-closed")
		if deleteAllClosed {
			storage.DeleteAllClosedIssues()
			fmt.Println("Successfully deleted all closed issues")
			return
		}
		err := storage.DeleteIssue(stringToInt(args[0]))
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Bool("all-closed", false, "Used to delete all closed issues")
}

// TODO: Refactor into utilities package
func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
