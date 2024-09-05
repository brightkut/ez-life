package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var rootCmd = &cobra.Command{
	Use:   "ez",
	Short: "ez is a CLI tool that simplifies various tasks to streamline your computer setup.",
	Long:  "ez is a CLI tool designed to streamline a variety of tasks, simplifying the process of setting up and managing your computer.",
	Run: func(cmd *cobra.Command, args []string) {
		// running help when type command `ez`
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing ez '%s'\n", err)
		os.Exit(1)
	}
}

func RunShScript(path string) {
	c := exec.Command("/bin/bash", path)

	// Pipe the output and error to Stdout and Stderr
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	// Run the command
	if err := c.Run(); err != nil {
		fmt.Printf("Could not run command: %v\n", err)
	}
}
