/*
Copyright © 2024 Parker Lee
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task or project",
	Long:  `Add a new task or project to the list of tasks or projects.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
