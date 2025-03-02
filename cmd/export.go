/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("export called")
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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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
