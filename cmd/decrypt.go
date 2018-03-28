// Copyright © 2017 Ibotta

package cmd

import (
	"fmt"

	"github.com/Ibotta/sopstool/execwrap"
	"github.com/Ibotta/sopstool/fileutil"
	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Aliases: []string{"d"},
	Use:     "decrypt [files ...]",
	Short:   "decrypt files",
	Long:    `Decrypt some or all files`,
	RunE:    DecryptCommand,
}

var allowFail bool

func init() {
	RootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().BoolVar(&allowFail, "allow-fail", false, "Do not fail if not all files can be decrypted")
}

// DecryptCommand decrypts files
func DecryptCommand(cmd *cobra.Command, args []string) error {
	initConfig()

	filesToDecrypt, err := fileutil.SomeOrAllFiles(args, sopsConfig.EncryptedFiles)
	if err != nil {
		return err
	}

	//decrypt all the files
	errCount := 0
	for _, f := range filesToDecrypt {
		err := execwrap.DecryptFile(f)
		if err != nil {
			if allowFail {
				errCount++
			} else {
				return err
			}
		}
	}

	if errCount == len(filesToDecrypt) {
		return fmt.Errorf("all files failed to decrypt")
	}

	return nil
}
