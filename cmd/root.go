package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "itc",
	Short: "itc is an all inclusive issue tracker for local projects",
	Long: `itc is an all inclusive local issue tracker. Data is stored in a SQLite Databse
named: ./.issues.db. itc currently only searches the root directory for this file. 
If not present, it will be created. Basic usage:

issue add -t "Fix login bug" -d "Users cannot log in with special characters" -p high -s open -a alice
issues list -s open
issue search "login bug"
issue edit 42 -t "Fix login error" -p critical -s in-progress
issue delete 42`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var title string
	rootCmd.PersistentFlags().StringVarP(&title, "title", "t", "", "Title for issue")

	var description string
	rootCmd.PersistentFlags().StringVarP(&description, "description", "d", "", "Description for issue")

	var priority string
	rootCmd.PersistentFlags().StringVarP(&priority, "priority", "p", "low", "Priority for issue")

	var status string
	rootCmd.PersistentFlags().StringVarP(&status, "status", "s", "open", "Status for issue")

	var user string
	rootCmd.PersistentFlags().StringVarP(&user, "user", "a", "", "User assigned to issue")
}
