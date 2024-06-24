/*
Copyright © 2024 Parker Lee
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Long: "Update a task from the list of tasks.",
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
