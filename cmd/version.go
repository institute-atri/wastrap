/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current installed version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Wastrap version: v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
