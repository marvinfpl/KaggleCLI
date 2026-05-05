package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xyproto/unzip"
	"os/exec"
	"fmt"
)


var listCmd = &cobra.Command{
	Use: "list",
	Short: "Returns a list of datasets",
	RunE: func (cmd *cobra.Command, args[] string) error {
		if len(args) != 1 {
			return fmt.Errorf("Unexpected number of arguments: got %d, want 1", len(args))
		}
		output, err := exec.Command("kaggle", "datasets", "list", args[0]).CombinedOutput()
		if err != nil {
			return fmt.Errorf("Invalid command: %v", err)
		}
		fmt.Println((string)(output))
		return nil
	},
}

var downloadCommand = &cobra.Command{
	Use: "download",
	Short: "Download the dataset corresponding to the given id",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Unexpected number of arguments: got %d, want 1", len(args))
		}
		_, err := exec.Command("kaggle", "datasets", "download", "-d", args[0]).CombinedOutput()
		if err != nil {
			return fmt.Errorf("Invalid command: %v", err)
		}
		return nil
	},
}

var unzipCmd = &cobra.Command{
	Use: "unzip",
	Short: "Unzip a given file, in a given destination",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Unexpected number of arguments: got %d, want 1", len(args))
		}
		cwd, err := exec.Command("pwd").CombinedOutput()
		if err != nil {
			return fmt.Errorf("Could not get current working directory: %v", err)
		}
		if err = unzip.Extract(args[0], (string)(cwd)); err != nil {
			return fmt.Errorf("Failed to unzip the dataset: %v", err)
		}
		return nil
	},
}