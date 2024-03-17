package cmd

import (
	"github.com/institute-atri/glogger"
	"github.com/institute-atri/wastrap/internal/banner"
	"github.com/spf13/cobra"
)

// Flags
var (
	url     string
	update  bool
	version bool
)

// rootCmd is the root command for the wastrap tool.
var rootCmd = &cobra.Command{
	Use:   "wastrap",
	Short: "WASTRAP is a web exploration tool focused on the WordPress application.",
	Long: `Being fast and simple, it has functionality that an information 
security professional cannot fail to have in their arsenal. 

With vulnerability analysis in themes, plugins and others, it 
also has brute-force in directories, users and much more (see 
more features in the official documentation).`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute is the main function of the package cmd
func Execute() {
	banner.Show()

	err := rootCmd.Execute()
	glogger.ErrorHandling(err)
}

// init sets up the command-line flags for the root command.
func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "Set the URL target")
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "Show the version of the tool")
	rootCmd.Flags().BoolVarP(&update, "update", "", false, "Update the tool")
}
