package cli

import (
	"os"
	"os/signal"

	"github.com/institute-atri/glogger"
	"github.com/institute-atri/wastrap/cli/cmd"
	"github.com/spf13/cobra"
)

func signalExit() {
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, os.Interrupt)

	<-sc

	glogger.Fatal("Your press CTRL + C")
}

var root = &cobra.Command{
	Use:   "wastrap",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: cmd.RootRun,
}

func Execute() {
	_ = root.Execute()
}

func init() {

	root.PersistentFlags().StringP("url", "u", "", "target url")
	root.PersistentFlags().BoolP("version", "V", false, "version")
	root.PersistentFlags().BoolP("update", "U", false, "update")
	root.MarkPersistentFlagRequired("url")
}
