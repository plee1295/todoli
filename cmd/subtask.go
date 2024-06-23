/*
Copyright © 2024 Parker Lee
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subtaskCmd = &cobra.Command{
	Use:   "subtask",
	Short: "Add a subtask",
	Long: "Add a subtask to a task.",
	Run: addSubtask,
}

func init() {
	addCmd.AddCommand(subtaskCmd)
}

func addSubtask(cmd *cobra.Command, args []string) {
	fmt.Println("addSubtask called")
}
