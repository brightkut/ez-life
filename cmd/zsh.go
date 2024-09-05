package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var zshCmd = &cobra.Command{
	Use:     "zsh",
	Aliases: []string{"zsh"},
	Short:   "install zsh terminal",
	Long:    "command for install zsh shell terminal",
	Args:    nil,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Trigger Zsh cmd\n")

		// Execute the shell script
		RunShScript("./cmd/sh/zsh.sh")
	},
}

func init() {
	rootCmd.AddCommand(zshCmd)
}
