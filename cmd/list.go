// Copyright © 2017 Ibotta

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list files under management",
	Long:  `Print a list of files under management`,
	Args:  cobra.NoArgs,
	Run:   ListCommand,
}

func init() {
	RootCmd.AddCommand(listCmd)

	//TODO flags format output (json, newline, yaml)?
}

// ListCommand the command for the list command
func ListCommand(cmd *cobra.Command, args []string) {
	for _, fn := range sopsConfig.EncryptedFiles {
		fmt.Println(fn)
	}
}
