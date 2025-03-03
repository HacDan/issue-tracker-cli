package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hacdan/issue-tracker-cli/storage"
	"github.com/hacdan/issue-tracker-cli/types"
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

		fileType, _ := cmd.Flags().GetString("filetype")
		if cmd.Flags().Changed("filetype") {
			if fileType == "json" {
				fmt.Println(issuesToJson(issues))
			}
			if fileType == "csv" {
				fmt.Println(issuesToCSV(issues))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringP("filetype", "f", "json", "Export file type")
}

// TODO: Refactor to utilities package
func issuesToJson(issues []types.Issue) string {
	sIssues, err := json.Marshal(issues)
	if err != nil {
		fmt.Println("Error marshaling json: ", err)
	}
	return string(sIssues)
}

func issuesToCSV(issues []types.Issue) string {
	header := []string{"id, title, description, priority, status, user"}
	data := [][]string{}
	data = append(data, header)
	for _, issue := range issues {
		row := []string{fmt.Sprintf("%d", issue.Id), issue.Title, issue.Description, issue.Priority.ToString(), issue.Status.ToString(), issue.User}
		data = append(data, row)
	}

	var result []string
	for _, row := range data {
		result = append(result, strings.Join(row, ","))
	}
	return strings.Join(result, strings.Join(result, "\n"))
}
