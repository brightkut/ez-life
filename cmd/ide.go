package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ideCmd = &cobra.Command{
	Use:     "ide",
	Aliases: []string{"ide"},
	Short:   "install Intelij",
	Long:    "command for install intelij with specific version and directory",
	Args:    nil,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Trigger IDE cmd\n")

		// Execute the shell script
		RunShScript("./cmd/sh/ide.sh")
	},
}

func init() {
	rootCmd.AddCommand(ideCmd)
}
