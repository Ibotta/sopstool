// Copyright © 2017 Ibotta

package cmd

import (
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
	decryptCmd.Flags().BoolVar(&allowFail, "allow-fail", false, "Do not fail if files can not be decrypted")
}

// DecryptCommand decrypts files
func DecryptCommand(_ *cobra.Command, args []string) error {
	initConfig()

	filesToDecrypt, err := fileutil.SomeOrAllFiles(args, sopsConfig.EncryptedFiles)
	if err != nil {
		return err
	}

	//decrypt all the files
	for _, f := range filesToDecrypt {
		err := encrypter.DecryptFile(f)
		if err != nil && !allowFail {
			return err
		}
	}

	return nil
}
