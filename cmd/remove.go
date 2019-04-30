// Copyright © 2017 Ibotta

package cmd

import (
	"fmt"

	"github.com/Ibotta/sopstool/fileutil"
	"github.com/Ibotta/sopstool/sopsyaml"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Aliases: []string{"rm"},
	Use:     "remove [files ...]",
	Short:   "remove file from the encryption list",
	Args:    cobra.MinimumNArgs(1),
	Long:    `Remove files to the list of files managed by sopstool`,
	RunE:    RemoveCommand,
}

var deleteFiles bool

func init() {
	RootCmd.AddCommand(removeCmd)

	removeCmd.Flags().BoolVarP(&deleteFiles, "delete", "d", false, "Also delete the file")
}

// RemoveCommand the command for the add command
func RemoveCommand(cmd *cobra.Command, args []string) error {
	initConfig()

	for _, fileArg := range args {
		fn := fileutil.NormalizeToPlaintextFile(fileArg)

		i := fileutil.ListIndexOf(sopsConfig.EncryptedFiles, fn)
		if i < 0 {
			return fmt.Errorf("File not found: %s", fn)
		}

		//splice file out of list
		sopsConfig.EncryptedFiles = append(sopsConfig.EncryptedFiles[:i], sopsConfig.EncryptedFiles[i+1:]...)

		if deleteFiles {
			err := encrypter.RemoveFile(fn)
			if err != nil {
				return err
			}
			err = encrypter.RemoveCryptFile(fn)
			if err != nil {
				return err
			}
		}

		fmt.Println("removed file from list:", fn)
	}

	err := sopsyaml.WriteEncryptFilesToDisk(sopsConfig.Path, sopsConfig.Tree, sopsConfig.EncryptedFiles)
	if err != nil {
		return err
	}

	return nil
}
