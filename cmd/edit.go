// Copyright © 2017 Ibotta

package cmd

import (
	"fmt"

	"github.com/Ibotta/sopstool/execwrap"
	"github.com/Ibotta/sopstool/fileutil"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit [file]",
	Long:  `edit a file in $EDITOR, reencrypting after completed`,
	RunE:  EditCommand,
}

func init() {
	RootCmd.AddCommand(editCmd)
}

// EditCommand edit a file
func EditCommand(cmd *cobra.Command, args []string) error {
	initConfig()

	fn := fileutil.NormalizeToPlaintextFile(args[0])
	if fileutil.ListIndexOf(sopsConfig.EncryptedFiles, fn) < 0 {
		return fmt.Errorf("File not found: %s", fn)
	}

	err := execwrap.EditFile(fn)
	if err != nil {
		return err
	}

	return nil
}
