package cmd

import (
	"github.com/spf13/cobra"
	"kaggle/internal/api"
	"fmt"
)

var downloadCmd = &cobra.Command{
	Use: "load",
	Short: "Download the given dataset from kaggle.com",
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]
		output := "dataset.zip"
		
		err := api.DownloadDataset(url, output)
		if err != nil {
			return err
		}
		fmt.Println("Dataset downloaded...")
		return nil
	},
}