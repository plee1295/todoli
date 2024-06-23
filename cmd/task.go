/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/alexeyco/simpletable"
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

var listTaskCmd = &cobra.Command{
	Use:   "task",
	Short: "List tasks",
	Long:  "List all tasks.",
	Run:   listTasks,
}

func init() {
	addCmd.AddCommand(addTaskCmd)
	addTaskCmd.Flags().StringP("project", "p", "", "Project name")

	deleteCmd.AddCommand(deleteTaskCmd)

	listCmd.AddCommand(listTaskCmd)
}

func addTask(cmd *cobra.Command, args []string) {
	task := types.Task{
		Status:    types.Open,
		CreatedAt: time.Now(),
	}

	if len(args) == 1 {
		task.Name = args[0]
	}

	if project, _ := cmd.Flags().GetString("project"); project != "" {
		var projects []types.Project
		if err := utils.ReadFromJSON(".projects.json", &projects); err != nil {
			fmt.Println("Error loading projects:", err)
			return
		}

		for _, p := range projects {
			if p.Name == project {
				task.ProjectID = p.ID
				break
			}
		}
	}

	task.Name, _ = utils.ReadInput("Please enter a task name", task.Name)
	task.Description, _ = utils.ReadInput("Please enter a task description", "")

	// Load existing tasks
	var tasks []types.Task
	if err := utils.ReadFromJSON(".tasks.json", &tasks); err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	task.ID = len(tasks) + 1

	// Append the new task
	tasks = append(tasks, task)

	// Save the updated list of tasks
	if err := utils.WriteToJSON(".tasks.json", tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("\nTask successfully added!")
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

func listTasks(cmd *cobra.Command, args []string) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "Status"},
			{Align: simpletable.AlignCenter, Text: "Project"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	var tasks []types.Task
	if err := utils.ReadFromJSON(".tasks.json", &tasks); err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	for _, item := range tasks {
		name := utils.Blue(item.Name)
		status := utils.Blue("no")

		switch {
		case item.Status == types.Open:
			name = item.Name
			status = "Open"
		case item.Status == types.InProgress:
			name = utils.Blue(item.Name)
			status = utils.Blue("In Progress")
		case item.Status == types.Completed:
			name = utils.Green(fmt.Sprintf("%s ✓", item.Name))
			status = utils.Green("Completed")
		}

		var projects []types.Project
		if err := utils.ReadFromJSON(".projects.json", &projects); err != nil {
			fmt.Println("Error loading projects:", err)
			return
		}

		projectName := "None"
		for _, p := range projects {
			if p.ID == item.ProjectID {
				projectName = p.Name
				break
			}
		}

		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", item.ID)},
			{Text: name},
			{Text: status},
			{Text: projectName},
			{Text: item.CreatedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: fmt.Sprintf("Total: %d", len(tasks))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}
