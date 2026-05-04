package cmd

import (
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use: "load",
	Short: "Download the given dataset from kaggle.com",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}