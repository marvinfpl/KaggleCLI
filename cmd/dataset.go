package cmd

import (
	"github.com/spf13/cobra"
	"os/exec"
	"fmt"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "Returns a list of datasets",
	RunE: func (cmd *cobra.Command, args[] string) error {
		output, err := exec.Command("kaggle", "datasets", "list", args[0]).CombinedOutput()
		if err != nil {
			return fmt.Errorf("Invalid command: %v", err)
		}
		fmt.Println((string)(output))
		return nil
	},
}