/*
Copyright © 2024 Parker Lee
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/plee1295/todoli/types"
	"github.com/plee1295/todoli/utils"
	"github.com/spf13/cobra"
)

func init() {
	commands := &types.Commands{
		Add: &cobra.Command{
			Use:   "subtask",
			Short: "Add a new subtask",
			Long:  "Add a new subtask to the list of subtasks.",
			Run:   addSubtask,
		},
		Delete: nil,
		View:   nil,
		List:   nil,
	}

	addCmd.AddCommand(commands.Add)
}

func addSubtask(cmd *cobra.Command, args []string) {
	subtask := types.Task{
		Status:    types.Open,
		CreatedAt: time.Now(),
	}

	id, err := utils.GenerateID()
	if err != nil {
		fmt.Println("Error generating ID:", err)
		return
	}

	subtask.ID = id

	if len(args) == 1 {
		subtask.Name = args[0]
	} else {
		subtask.Name, _ = utils.ReadInput("Enter subtask name", subtask.Name)
	}

	subtask.Description, _ = utils.ReadInput("Enter subtask description", subtask.Description)

	var allTasks []types.Task
	if err := utils.ReadFromJSON(".tasks.json", &allTasks); err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	var parentTasks []string
	for _, task := range allTasks {
		if task.ParentID == "" {
			parentTasks = append(parentTasks, task.Name)
		}
	}

	if len(parentTasks) == 0 {
		fmt.Println("No tasks to add subtask to")
		return
	}

	parentTaskChoice, _ := utils.ReadMultipleChoice("Choose a parent task", parentTasks)
	for _, task := range allTasks {
		if task.Name == parentTaskChoice {
			subtask.ParentID = task.ID
			subtask.ProjectID = task.ProjectID
		}
	}

	statusChoice, _ := utils.ReadMultipleChoice("Subtask status", []string{"Open", "In Progress", "Completed"})
	switch statusChoice {
	case "Open":
		subtask.Status = types.Open
	case "In Progress":
		subtask.Status = types.InProgress
	case "Completed":
		subtask.Status = types.Completed
	}

	priorityChoice, _ := utils.ReadMultipleChoice("Subtask priority", []string{"Critical", "High", "Medium", "Low"})
	switch priorityChoice {
	case "Critical":
		subtask.Priority = types.Critical
	case "High":
		subtask.Priority = types.High
	case "Medium":
		subtask.Priority = types.Medium
	case "Low":
		subtask.Priority = types.Low
	}

	allTasks = append(allTasks, subtask)

	if err := utils.WriteToJSON(".tasks.json", allTasks); err != nil {
		fmt.Println("Error writing tasks:", err)
		return
	}

	fmt.Println("Subtask added successfully")
}
