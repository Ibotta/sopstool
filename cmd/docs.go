// Copyright © 2017 Ibotta
// https://github.com/spf13/cobra/blob/master/doc/md_docs.md

package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:    "docs",
	Short:  "get docs",
	Long:   `get docs for this`,
	Hidden: true,
	RunE:   DocsCommand,
}

var docsFormat string

func init() {
	RootCmd.AddCommand(docsCmd)

	// TODO also support man
	docsCmd.Flags().StringVarP(&docsFormat, "format", "f", "md", "format")
}

// DocsCommand the command for the add command
func DocsCommand(cmd *cobra.Command, args []string) error {
	err := doc.GenMarkdownTree(RootCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
