package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var hyperCmd = &cobra.Command{
	Use:     "hyper",
	Aliases: []string{"hyper"},
	Short:   "install hyper",
	Long:    "command for install hyper",
	Args:    nil,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Trigger Hyper cmd\n")

		// Execute the shell script
		RunShScript("./cmd/sh/hyper.sh")
	},
}

func init() {
	rootCmd.AddCommand(hyperCmd)
}
