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
	Use:   "cat [file]",
	Short: "print a file to stdout",
	Long:  `Decrypt a file and print it to stdout`,
	Args:  cobra.ExactArgs(1),
	RunE:  CatCommand,
}

func init() {
	RootCmd.AddCommand(catCmd)
}

// CatCommand prints a file to stdout
func CatCommand(cmd *cobra.Command, args []string) error {
	fn := fileutil.NormalizeToPlaintextFile(args[0])
	if fileutil.ListIndexOf(sopsConfig.EncryptedFiles, fn) < 0 {
		return fmt.Errorf("File not found: %s", fn)
	}

	err := execwrap.DecryptFilePrint(fn)
	if err != nil {
		return err
	}

	return nil
}
