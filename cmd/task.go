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

func init() {
	commands := &types.Commands{
		Add: &cobra.Command{
			Use:   "task",
			Short: "Add a new task",
			Long:  "Add a new task to the list of tasks.",
			Run:   addTask,
		},
		Delete: &cobra.Command{
			Use:   "task",
			Short: "Delete a task",
			Long:  "Delete a task from the list of tasks.",
			Run:   deleteTask,
		},
		Update: &cobra.Command{
			Use:   "task",
			Short: "Update a task",
			Long:  "Update a task from the list of tasks.",
			Run:   updateTask,
		},
		View: &cobra.Command{
			Use:   "task",
			Short: "View a task",
			Long:  "View a task from the list of tasks.",
			Run:   viewTask,
		},
		List: &cobra.Command{
			Use:   "tasks",
			Short: "List tasks",
			Long:  "List all tasks.",
			Run:   listTasks,
		},
	}

	addCmd.AddCommand(commands.Add)
	commands.Add.Flags().StringP("project", "p", "", "Project name")

	viewCmd.AddCommand(commands.View)
	deleteCmd.AddCommand(commands.Delete)
	updateCmd.AddCommand(commands.Update)
	listCmd.AddCommand(commands.List)
}

func addTask(cmd *cobra.Command, args []string) {
	task := types.Task{
		Status:    types.Open,
		CreatedAt: time.Now(),
	}

	id, err := utils.GenerateID()
	if err != nil {
		fmt.Println("Error generating ID:", err)
		return
	}

	task.ID = id

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

	task.Name, _ = utils.ReadInput("Task name", task.Name)
	task.Description, _ = utils.ReadInput("Task description", "")

	statusChoice, _ := utils.ReadMultipleChoice("Task status", []string{"Open", "In Progress", "Completed"})
	switch statusChoice {
	case "Open":
		task.Status = types.Open
	case "In Progress":
		task.Status = types.InProgress
	case "Completed":
		task.Status = types.Completed
	}

	priorityChoice, _ := utils.ReadMultipleChoice("Task priority", []string{"Critical", "High", "Medium", "Low"})
	switch priorityChoice {
	case "Critical":
		task.Priority = types.Critical
	case "High":
		task.Priority = types.High
	case "Medium":
		task.Priority = types.Medium
	case "Low":
		task.Priority = types.Low
	}

	// TODO: Allow for multiple labels to be selected
	var labels []types.Label
	if err := utils.ReadFromJSON(".labels.json", &labels); err != nil {
		fmt.Println("Error loading labels:", err)
		return
	}

	var labelStrings []string
	for _, label := range labels {
		labelStrings = append(labelStrings, label.String())
	}

	labelChoice, _ := utils.ReadMultipleChoice("Task labels", labelStrings)
	task.Labels = append(task.Labels, types.Label(labelChoice))

	var projects []types.Project
	if err := utils.ReadFromJSON(".projects.json", &projects); err != nil {
		fmt.Println("Error loading projects:", err)
		return
	}

	var projectStrings []string
	for _, project := range projects {
		projectStrings = append(projectStrings, project.String())
	}

	projectChoice, _ := utils.ReadMultipleChoice("Task project", projectStrings)
	for _, project := range projects {
		if project.String() == projectChoice {
			task.ProjectID = project.ID
			break
		}
	}

	var tasks []types.Task
	if err := utils.ReadFromJSON(".tasks.json", &tasks); err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	tasks = append(tasks, task)

	if err := utils.WriteToJSON(".tasks.json", tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("\nTask successfully added!")
}

func deleteTask(cmd *cobra.Command, args []string) {
	var tasks []types.Task
	if err := utils.ReadFromJSON(".tasks.json", &tasks); err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	// Loop over tasks in reverse order to avoid index out of range errors
	var task types.Task
	for i := len(tasks) - 1; i >= 0; i-- {
		t := tasks[i]
		// Remove task and all subtasks
		if t.ID == args[0] || t.ParentID == args[0] {
			task = t
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}

	if err := utils.WriteToJSON(".tasks.json", tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("\nTask successfully deleted!", task)
}

func updateTask(cmd *cobra.Command, args []string) {
	var tasks []types.Task
	if err := utils.ReadFromJSON(".tasks.json", &tasks); err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	id := args[0]
	var task types.Task
	for _, t := range tasks {
		if t.ID == id {
			task = t
		}
	}

	// Todo - Create loop to update multiple properties
	fmt.Println("Why property would you like to update?")
	statusChoice, _ := utils.ReadMultipleChoice("Properties", []string{"Name", "Description", "Project", "Parent Task", "Status", "Priority", "Labels"})
	switch statusChoice {
	case "Name":
		task.Name, _ = utils.ReadInput("Task name", task.Name)

	case "Description":
		task.Description, _ = utils.ReadInput("Task description", task.Description)

	case "Project":
		var projects []types.Project
		if err := utils.ReadFromJSON(".projects.json", &projects); err != nil {
			fmt.Println("Error loading projects:", err)
			return
		}

		var projectStrings []string
		for _, project := range projects {
			projectStrings = append(projectStrings, project.String())
		}

		projectChoice, _ := utils.ReadMultipleChoice("Task project", projectStrings)
		for _, project := range projects {
			if project.String() == projectChoice {
				task.ProjectID = project.ID
				break
			}
		}

	case "Parent Task":
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
		for _, t := range allTasks {
			if t.Name == parentTaskChoice {
				task.ParentID = t.ID
			}
		}

	case "Status":
		statusChoice, _ := utils.ReadMultipleChoice("Task status", []string{"Open", "In Progress", "Completed"})
		switch statusChoice {
		case "Open":
			task.Status = types.Open
		case "In Progress":
			task.Status = types.InProgress
		case "Completed":
			task.Status = types.Completed
		}

	case "Priority":
		priorityChoice, _ := utils.ReadMultipleChoice("Task priority", []string{"Critical", "High", "Medium", "Low"})
		switch priorityChoice {
		case "Critical":
			task.Priority = types.Critical
		case "High":
			task.Priority = types.High
		case "Medium":
			task.Priority = types.Medium
		case "Low":
			task.Priority = types.Low
		}

	case "Labels":
		var labels []types.Label
		if err := utils.ReadFromJSON(".labels.json", &labels); err != nil {
			fmt.Println("Error loading labels:", err)
			return
		}

		var labelStrings []string
		for _, label := range labels {
			labelStrings = append(labelStrings, label.String())
		}

		labelChoice, _ := utils.ReadMultipleChoice("Task labels", labelStrings)
		task.Labels = append(task.Labels, types.Label(labelChoice))
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i] = task
		}
	}

	if err := utils.WriteToJSON(".tasks.json", tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("\nTask successfully updated!")
}

func viewTask(cmd *cobra.Command, args []string) {
	var tasks []types.Task
	if err := utils.ReadFromJSON(".tasks.json", &tasks); err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Subtask"},
			{Align: simpletable.AlignCenter, Text: "Status"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	taskFound := false

	for _, task := range tasks {
		if task.ID == args[0] {
			taskFound = true

			fmt.Println("ID:", task.ID)
			fmt.Println("Name:", task.Name)
			fmt.Println("Description:", task.Description)
			fmt.Println("Project ID:", task.ProjectID)
			fmt.Println("Parent ID:", task.ParentID)
			fmt.Println("Status:", task.Status)
			fmt.Println("Priority:", task.Priority)
			fmt.Println("Labels:", task.Labels)
			fmt.Println("Created At:", task.CreatedAt)
			fmt.Println("Due At:", task.DueAt)
			fmt.Println("Completed At:", task.CompletedAt)
		}

		if task.ParentID == args[0] {
			cells = append(cells, []*simpletable.Cell{
				{Text: task.ID},
				{Text: task.Name},
				{Text: task.Status.String()},
				{Text: task.CreatedAt.Format(time.RFC822)},
			})
		}
	}

	if !taskFound {
		fmt.Printf("Task with ID %s not found\n", args[0])
		return
	}

	table.Body = &simpletable.Body{Cells: cells}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
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
			{Align: simpletable.AlignCenter, Text: "Subtasks"},
		},
	}

	var cells [][]*simpletable.Cell

	var tasks []types.Task
	if err := utils.ReadFromJSON(".tasks.json", &tasks); err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	type SubtaskInfo struct {
		Total     int
		Completed int
	}

	subtaskCount := make(map[string]SubtaskInfo)
	for _, task := range tasks {
		if task.ParentID != "" {
			info := subtaskCount[task.ParentID]
			info.Total++
			if task.Status == types.Completed {
				info.Completed++
			}
			subtaskCount[task.ParentID] = info
		} else {
			subtaskCount[task.ID] = SubtaskInfo{Total: 0, Completed: 0}
		}
	}

	parentTaskCount := 0
	for _, task := range tasks {
		if task.ParentID == "" {
			parentTaskCount++
			name := utils.Blue(task.Name)
			status := utils.Blue("no")

			switch {
			case task.Status == types.Open:
				name = task.Name
				status = "Open"
			case task.Status == types.InProgress:
				name = utils.Blue(task.Name)
				status = utils.Blue("In Progress")
			case task.Status == types.Completed:
				name = utils.Green(fmt.Sprintf("%s ✓", task.Name))
				status = utils.Green("Completed")
			}

			var projects []types.Project
			if err := utils.ReadFromJSON(".projects.json", &projects); err != nil {
				fmt.Println("Error loading projects:", err)
				return
			}

			projectName := "-"
			for _, p := range projects {
				if p.ID == task.ProjectID {
					projectName = p.Name
					break
				}
			}

			completed := subtaskCount[task.ID].Completed
			total := subtaskCount[task.ID].Total
			percent := 0
			if total > 0 {
				percent = int(float64(completed)/float64(total)*100 + 0.5)
			}

			var subtaskStr string
			if total == 0 {
				subtaskStr = "-"
			} else {
				subtaskStr = fmt.Sprintf("%d/%d (%d%%)", completed, total, percent)
			}

			cells = append(cells, []*simpletable.Cell{
				{Text: task.ID},
				{Text: name},
				{Text: status},
				{Text: projectName},
				{Text: task.CreatedAt.Format(time.RFC822)},
				{Text: subtaskStr},
			})
		}
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignRight, Span: 6, Text: fmt.Sprintf("Total: %d", parentTaskCount)},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}
