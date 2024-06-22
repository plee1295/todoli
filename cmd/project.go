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

type Project struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Tasks       []Task    `json:"tasks"`
	CreatedAt   time.Time `json:"created_at"`
}

var addProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Add a new project",
	Long:  "Add a new project to the list of projects.",
	Run:   addProject,
}

var deleteProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Delete a project",
	Long:  "Delete a project from the list of projects.",
	Run:   deleteProject,
}

func init() {
	addCmd.AddCommand(addProjectCmd)
	deleteCmd.AddCommand(deleteProjectCmd)
}

func addProject(cmd *cobra.Command, args []string) {
	project := Project{
		Tasks:     []Task{},
		CreatedAt: time.Now(),
	}

	if len(args) == 1 {
		project.Name = args[0]
	}

	project.Name, _ = utils.ReadInput("Please enter a project name", project.Name)
	project.Description, _ = utils.ReadInput("Please enter a project description", "")

	// Load existing projects
	var projects []Project
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

	fmt.Println("\nProject successfully added!", project)
}

func deleteProject(cmd *cobra.Command, args []string) {
	// Load existing projects
	var projects []Project
	if err := utils.ReadFromJSON(".projects.json", &projects); err != nil {
		fmt.Println("Error loading projects:", err)
		return
	}

	// Find the project to delete
	var project Project
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
