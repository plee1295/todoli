/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/plee1295/todoli/utils"
	"github.com/spf13/cobra"
)

type Task struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Project     string    `json:"project"`
	Status      string    `json:"status"`
	Priority    int       `json:"priority"`
	Labels      []string  `json:"labels"`
	CreatedAt   time.Time `json:"created_at"`
	DueAt       time.Time `json:"due_at"`
	CompletedAt time.Time `json:"completed_at"`
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Add a new task",
	Long:  "Add a new task to the list of tasks.",
	Run:   addTask,
}

func init() {
	addCmd.AddCommand(taskCmd)

	taskCmd.Flags().StringP("project", "p", "", "Project name")
}

func addTask(cmd *cobra.Command, args []string) {
	task := Task{
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
	var tasks []Task
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
