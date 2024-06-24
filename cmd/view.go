/*
Copyright © 2024 Parker Lee
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View a task",
	Long: "View a task from the list of tasks.",
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
