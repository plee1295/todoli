/*
Copyright © 2024 Parker Lee
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task or project",
	Long: "Delete a task or project from the list of tasks or projects.",
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
