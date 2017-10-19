// Copyright © 2017 Ibotta

package cmd

import (
	"github.com/Ibotta/go-commons/sopstool/execwrap"
	"github.com/Ibotta/go-commons/sopstool/fileutil"
	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt [files ...]",
	Short: "decrypt files",
	Long:  `Decrypt some or all files`,
	RunE:  DecryptCommand,
}

func init() {
	RootCmd.AddCommand(decryptCmd)
}

// DecryptCommand decrypts files
func DecryptCommand(cmd *cobra.Command, args []string) error {
	filesToDecrypt, err := fileutil.SomeOrAllFiles(args, sopsConfig.EncryptedFiles)
	if err != nil {
		return err
	}

	//decrypt all the files
	for _, f := range filesToDecrypt {
		err := execwrap.DecryptFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
