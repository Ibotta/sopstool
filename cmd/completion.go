// Copyright © 2017 Ibotta
// https://github.com/spf13/cobra/blob/main/bash_completions.md.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "get completions",
	Long:  `get command completions`,
	RunE:  CompletionCommand,
}

var shellType string

func init() {
	RootCmd.AddCommand(completionCmd)

	// TODO also support man
	completionCmd.Flags().StringVar(&shellType, "sh", "bash", "format")
}

// CompletionCommand the command for the add command
func CompletionCommand(_ *cobra.Command, _ []string) error {
	switch shellType {
	case "bash":
		return RootCmd.GenBashCompletion(os.Stdout)
	case "zsh":
		return RootCmd.GenZshCompletion(os.Stdout)
	default:
		return fmt.Errorf("invalid shell type")
	}
}
