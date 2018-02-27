// Copyright © 2017 Ibotta

package cmd

import (
	"fmt"

	"github.com/Ibotta/sopstool/execwrap"
	"github.com/Ibotta/sopstool/fileutil"
	"github.com/Ibotta/sopstool/sopsyaml"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Aliases: []string{"a", "encrypt"},
	Use:     "add [files ...]",
	Short:   "add file to the encryption list",
	Long:    `Add files to the list of files managed by sopstool`,
	Args:    cobra.MinimumNArgs(1),
	RunE:    AddCommand,
}

var noEncrypt bool

func init() {
	RootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolVarP(&noEncrypt, "no-encrypt", "n", false, "Do not encrypt the file after adding")
}

// AddCommand the command for the add command
func AddCommand(cmd *cobra.Command, args []string) error {
	for _, fileArg := range args {
		fn := fileutil.NormalizeToPlaintextFile(fileArg)

		if fileutil.ListIndexOf(sopsConfig.EncryptedFiles, fn) >= 0 {
			fmt.Println("Already exists", fn)
			return nil
		}

		// add file to list
		sopsConfig.EncryptedFiles = append(sopsConfig.EncryptedFiles, fn)
		fmt.Println("added file to list:", fn)

		//if the file exists, encrypt it
		if !noEncrypt {
			err := execwrap.EncryptFile(fn)
			if err != nil {
				return err
			}
		}
	}

	err := sopsyaml.WriteEncryptFilesToDisk(sopsConfig.Path, sopsConfig.Tree, sopsConfig.EncryptedFiles)
	if err != nil {
		return err
	}

	return nil
}
