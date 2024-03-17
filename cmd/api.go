package cmd

import (
	"github.com/spf13/cobra"
)

// apiCmd is the root command for the wastrap tool.
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "API Restful for the tool",
	Long:  `It is possible to use the tool's features through the API.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
	},
}

// init sets up the command-line flags for the api command.
func init() {
	rootCmd.AddCommand(apiCmd)
}
