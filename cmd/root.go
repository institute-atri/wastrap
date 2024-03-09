/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wastrap",
	Short: "-WordPress Security Scanner",
	Long: `-Wastrap is a comprehensive security scanner for WordPress websites.
-It scans for vulnerabilities and provides information to improve the overall security of your WordPress installations.
-The CLI supports various commands to perform specific tasks.

-For more information, visit: https://github.com/institute-atri/wastrap`,
	Run: func(cmd *cobra.Command, args []string) {
		showWelcomeMessage()
		_ = cmd.Help()
	},
}

func showWelcomeMessage() {
	fmt.Printf("\n🛡️ %s 🛡️\n", formatMessage("Welcome to Wastrap - WordPress Security Scanner", 1))
	fmt.Println()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
