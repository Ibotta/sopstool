// Copyright © 2017 Ibotta

package cmd

import (
	"fmt"

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
var noClean bool

func init() {
	RootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolVarP(&noEncrypt, "no-encrypt", "n", false, "Do not encrypt the file after adding")
	addCmd.Flags().BoolVar(&noClean, "no-clean", false, "Do not clean up plaintext after encrypting")
}

// AddCommand the command for the add command
func AddCommand(_ *cobra.Command, args []string) error {
	initConfig()

	for _, fileArg := range args {
		fn := fileutil.NormalizeToPlaintextFile(fileArg)

		if fileutil.ListIndexOf(sopsConfig.EncryptedFiles, fn) >= 0 {
			fmt.Println("Already exists", fn)
			return nil
		}

		// add file to list
		sopsConfig.EncryptedFiles = append(sopsConfig.EncryptedFiles, fn)

		//if the file exists, encrypt it
		if !noEncrypt {
			err := encrypter.EncryptFile(fn)
			if err != nil {
				return err
			}
		}

		if !noClean {
			err := encrypter.RemoveFile(fn)
			if err != nil {
				return err
			}
		}

		err := sourceCodeManager.AddFileToIgnored(fn)
		if err != nil {
			return err
		}
		fmt.Println("added file to list:", fn)
	}

	err := sopsyaml.WriteEncryptFilesToDisk(sopsConfig.Path, sopsConfig.Tree, sopsConfig.EncryptedFiles)
	if err != nil {
		return err
	}

	fmt.Println("Files added")

	return nil
}
