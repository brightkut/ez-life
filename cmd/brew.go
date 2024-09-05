package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var brewCmd = &cobra.Command{
	Use:     "brew",
	Aliases: []string{"brew"},
	Short:   "install Homebrew",
	Long:    "command for install home brew",
	Args:    nil,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Trigger Home brew cmd\n")

		// Execute the shell script
		RunShScript("./cmd/sh/brew.sh")
	},
}

func init() {
	rootCmd.AddCommand(brewCmd)
}
