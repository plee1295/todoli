/*
Copyright © 2024 Parker Lee
*/
package cmd

import (
	"fmt"

	"github.com/plee1295/todoli/types"
	"github.com/plee1295/todoli/utils"
	"github.com/spf13/cobra"
)

func init() {
	commands := &types.Commands{
		Add: &cobra.Command{
			Use:   "label",
			Short: "Add a new label",
			Long:  "Add a new label to the list of labels.",
			Run:   addLabel,
		},
		Delete: nil,
		View: nil,
		List: nil,
	}

	addCmd.AddCommand(commands.Add)
}

func addLabel(cmd *cobra.Command, args []string) {
	var label types.Label

	if len(args) == 1 {
		label = types.Label(args[0])
	}

	var labels []types.Label
	if err := utils.ReadFromJSON(".labels.json", &labels); err != nil {
		fmt.Println("Error loading labels:", err)
		return
	}

	labels = append(labels, label)

	if err := utils.WriteToJSON(".labels.json", labels); err != nil {
		fmt.Println("Error saving labels:", err)
		return
	}

	fmt.Println("Label successfully added!")
}
