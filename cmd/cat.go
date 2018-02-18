// Copyright © 2017 Ibotta

package cmd

import (
	"fmt"

	"github.com/Ibotta/sopstool/execwrap"
	"github.com/Ibotta/sopstool/fileutil"

	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat [files ...]",
	Short: "print files to stdout",
	Long:  `Decrypt files and print to stdout`,
	Args:  cobra.MinimumNArgs(1),
	RunE:  CatCommand,
}

func init() {
	RootCmd.AddCommand(catCmd)
}

// CatCommand prints a file to stdout
func CatCommand(cmd *cobra.Command, args []string) error {
	for _, fileArg := range args {
		fn := fileutil.NormalizeToPlaintextFile(fileArg)
		if fileutil.ListIndexOf(sopsConfig.EncryptedFiles, fn) < 0 {
			return fmt.Errorf("File not found: %s", fn)
		}

		err := execwrap.DecryptFilePrint(fn)
		if err != nil {
			return err
		}
	}

	return nil
}
