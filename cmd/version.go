// Copyright © 2017 Ibotta

package cmd

import (
	"fmt"

	"github.com/Ibotta/go-commons/sopstool/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Aliases: []string{"v"},
	Use:     "version",
	Short:   "Program version information",
	Run:     VersionCommand,
}

var shortString bool

func init() {
	RootCmd.AddCommand(versionCmd)

	versionCmd.Flags().BoolVarP(&shortString, "short", "s", false, "Only print the short version tag")
}

// VersionCommand prints the version
func VersionCommand(cmd *cobra.Command, args []string) {
	if shortString {
		fmt.Println(version.Number)
	} else {
		fmt.Printf("%s v%s\n", RootCmd.Use, version.Long)
	}
}
