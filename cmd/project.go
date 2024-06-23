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
			Use:   "project",
			Short: "Add a new project",
			Long:  "Add a new project to the list of projects.",
			Run:   addProject,
		},
		Delete: &cobra.Command{
			Use:   "project",
			Short: "Delete a project",
			Long:  "Delete a project from the list of projects.",
			Run:   deleteProject,
		},
		List: &cobra.Command{
			Use:   "projects",
			Short: "List projects",
			Long:  "List all projects.",
			Run:   listProjects,
		},
	}

	addCmd.AddCommand(commands.Add)
	deleteCmd.AddCommand(commands.Delete)
	listCmd.AddCommand(commands.List)
}

func addProject(cmd *cobra.Command, args []string) {
	project := types.Project{
		CreatedAt: time.Now(),
	}

	id, err := utils.GenerateID()
	if err != nil {
		fmt.Println("Error generating ID:", err)
		return
	}

	project.ID = id

	if len(args) == 1 {
		project.Name = args[0]
	}

	project.Name, _ = utils.ReadInput("Please enter a project name", project.Name)
	project.Description, _ = utils.ReadInput("Please enter a project description", "")

	// Load existing projects
	var projects []types.Project
	if err := utils.ReadFromJSON(".projects.json", &projects); err != nil {
		fmt.Println("Error loading projects:", err)
		return
	}

	// Append the new project
	projects = append(projects, project)

	// Save the updated list of projects
	if err := utils.WriteToJSON(".projects.json", projects); err != nil {
		fmt.Println("Error saving projects:", err)
		return
	}

	fmt.Println("\nProject successfully added!")
}

func deleteProject(cmd *cobra.Command, args []string) {
	// Load existing projects
	var projects []types.Project
	if err := utils.ReadFromJSON(".projects.json", &projects); err != nil {
		fmt.Println("Error loading projects:", err)
		return
	}

	// Find the project to delete
	var project types.Project
	for i, t := range projects {
		if t.Name == args[0] {
			project = t
			projects = append(projects[:i], projects[i+1:]...)
			break
		}
	}

	// Save the updated list of projects
	if err := utils.WriteToJSON(".projects.json", projects); err != nil {
		fmt.Println("Error saving projects:", err)
		return
	}

	fmt.Println("\nProject successfully deleted!", project)
}

func listProjects(cmd *cobra.Command, args []string) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	var projects []types.Task
	if err := utils.ReadFromJSON(".projects.json", &projects); err != nil {
		fmt.Println("Error loading projects:", err)
		return
	}

	for _, project := range projects {
		cells = append(cells, []*simpletable.Cell{
			{Text: project.ID},
			{Text: project.Name},
			{Text: project.CreatedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignRight, Span: 3, Text: fmt.Sprintf("Total: %d", len(projects))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}
