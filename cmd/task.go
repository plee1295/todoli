/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/plee1295/todoli/types"
	"github.com/plee1295/todoli/utils"
	"github.com/spf13/cobra"
)

var addTaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Add a new task",
	Long:  "Add a new task to the list of tasks.",
	Run:   addTask,
}

var deleteTaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Delete a task",
	Long:  "Delete a task from the list of tasks.",
	Run:   deleteTask,
}

func init() {
	addCmd.AddCommand(addTaskCmd)
	addTaskCmd.Flags().StringP("project", "p", "", "Project name")

	deleteCmd.AddCommand(deleteTaskCmd)
}

func addTask(cmd *cobra.Command, args []string) {
	task := types.Task{
		CreatedAt: time.Now(),
	}

	if len(args) == 1 {
		task.Name = args[0]
	}

	if project, _ := cmd.Flags().GetString("project"); project != "" {
		task.Project = project
	}

	task.Name, _ = utils.ReadInput("Please enter a task name", task.Name)
	task.Description, _ = utils.ReadInput("Please enter a task description", "")

	// Load existing tasks
	var tasks []types.Task
	if err := utils.ReadFromJSON(".tasks.json", &tasks); err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	// Append the new task
	tasks = append(tasks, task)

	// Save the updated list of tasks
	if err := utils.WriteToJSON(".tasks.json", tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("\nTask successfully added!", task)
}

func deleteTask(cmd *cobra.Command, args []string) {
	// Load existing tasks
	var tasks []types.Task
	if err := utils.ReadFromJSON(".tasks.json", &tasks); err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	// Find the task to delete
	var task types.Task
	for i, t := range tasks {
		if t.Name == args[0] {
			task = t
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	// Save the updated list of tasks
	if err := utils.WriteToJSON(".tasks.json", tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("\nTask successfully deleted!", task)
}
