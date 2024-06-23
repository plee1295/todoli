/*
Copyright © 2024 Parker Lee
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks or projects",
	Long:  "List all tasks or projects.",
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
