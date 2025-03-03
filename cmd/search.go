package cmd

import (
	"fmt"

	"github.com/hacdan/issue-tracker-cli/storage"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search searches all issues in the database.",
	Long: `search is used to search all issues in the database. Can be used to search inside
the title or description. If not flag is specified, searches both fields.
Examples:

issue search "login bug"
issue search -t "UI glitch"       # Search by title  
issue search -d "performance"     # Search in description  `,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")

		fmt.Println(title, description)
		if len(args) > 0 {
			storage := storage.NewStorage()
			issues, err := storage.SearchIssuesByText(args[0])
			if err != nil {
				fmt.Println(err)
			}
			for _, issue := range issues {
				issue.Print()
			}
		}
		if cmd.Flags().Changed("title") {
			storage := storage.NewStorage()
			issues, err := storage.SearchIssuesByTitle(title)
			if err != nil {
				fmt.Println(err)
			}

			for _, issue := range issues {
				issue.Print()
			}
			return
		}
		if cmd.Flags().Changed("description") {
			storage := storage.NewStorage()
			issues, err := storage.SearchIssuesByDescription(description)
			if err != nil {
				fmt.Println(err)
			}
			for _, issue := range issues {
				issue.Print()
			}
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
