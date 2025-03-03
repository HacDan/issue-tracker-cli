package cmd

import (
	"fmt"

	"github.com/hacdan/issue-tracker-cli/storage"
	"github.com/hacdan/issue-tracker-cli/utils"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports all issues in the database",
	Long: `Takes all issues in the database, 
converts to the specified structured data format (JSON, CSV)
(defaults to JSON) and writes to stdout. Examples:

issue export -f json > issues.json
issue export -f csv > issues.csv`,
	Run: func(cmd *cobra.Command, args []string) {
		store := storage.NewStorage()
		issues, _ := store.GetIssues()

		fileType, err := cmd.Flags().GetString("filetype")
		if err != nil {
			fmt.Println(err)
		}
		if fileType == "json" {
			fmt.Println(utils.IssuesToJson(issues))
			return
		}
		if fileType == "csv" {
			fmt.Println(utils.IssuesToCSV(issues))
			return
		}
		fmt.Println(utils.IssuesToJson(issues))
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringP("filetype", "f", "json", "Export file type")
}
