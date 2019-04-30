// Copyright © 2017 Ibotta

package cmd

import (
	"github.com/Ibotta/sopstool/fileutil"
	"github.com/spf13/cobra"
)

// rotateCmd represents the rotate command
var rotateCmd = &cobra.Command{
	Use:   "rotate [files ...]",
	Short: "rotate keys on files",
	Long:  `Re-encrypt and rotate data the keys on some or all files`,
	RunE:  RotateCommand,
}

func init() {
	RootCmd.AddCommand(rotateCmd)
}

// RotateCommand Rotates up files
func RotateCommand(cmd *cobra.Command, args []string) error {
	initConfig()

	filesToRotate, err := fileutil.SomeOrAllFiles(args, sopsConfig.EncryptedFiles)
	if err != nil {
		return err
	}

	//Rotate all the files
	for _, f := range filesToRotate {
		err := encrypter.RotateFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
