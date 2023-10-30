// Copyright © 2017 Ibotta

package cmd

import (
	"fmt"
	"os"

	"github.com/Ibotta/sopstool/fileutil"
	"github.com/Ibotta/sopstool/sopsyaml"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Aliases: []string{"a", "e", "add", "encrypt"},
	Use:     "add [files ...], add",
	Short:   "add file to the encryption list and encrypt file. No argument encrypts all in sops config.",
	Long:    `Add files to the list of files managed by sopstool and encrypt them. If no argument is provided, encrypt everything provided in sops config`,
	RunE:    AddCommand,
}

var noEncrypt bool
var noClean bool
var forceOverwrite bool

func init() {
	RootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolVarP(&noEncrypt, "no-encrypt", "n", false, "Do not encrypt the file after adding")
	addCmd.Flags().BoolVar(&noClean, "no-clean", false, "Do not clean up plaintext after encrypting")
	addCmd.Flags().BoolVarP(&forceOverwrite, "force", "f", false, "Force overwriting of encrypted files if they already exist")
}

func AddCommand(_ *cobra.Command, args []string) error {
	initConfig()

	if len(args) == 0 && len(sopsConfig.EncryptedFiles) == 0 {
		fmt.Println("No files specified for encryption and no known encrypted files.")
		return nil
	}

	filesToEncrypt := args
	if len(args) == 0 {
		filesToEncrypt = sopsConfig.EncryptedFiles
	}

	encryptedCount := 0
	for _, fileArg := range filesToEncrypt {
		fn := fileutil.NormalizeToPlaintextFile(fileArg)


		// Check if plaintext file exists on disk
		if _, err := os.Stat(fn); os.IsNotExist(err) {
			fmt.Println("Plaintext version of file does not exist:", fn)
			continue
		}

		encryptedFilePath := fileutil.NormalizeToSopsFile(fn)

		// Check if encrypted file already exists on disk
		if _, err := os.Stat(encryptedFilePath); !os.IsNotExist(err) && !forceOverwrite {
			fmt.Printf("Encrypted version of file %s already exists on disk. Do you want to overwrite it? [y/N]: ", encryptedFilePath)
			var input string
			fmt.Scanln(&input)
			if input == "" || input == "y" || input == "Y" {
				// Continue with the overwriting process
			} else {
				fmt.Println("Skipping:", encryptedFilePath)
				continue
			}
		}

		// Check and add the file to the list if not already present
		if fileutil.ListIndexOf(sopsConfig.EncryptedFiles, fn) == -1 {
			sopsConfig.EncryptedFiles = append(sopsConfig.EncryptedFiles, fn)
			fmt.Println("Added file to encryption list:", fn)
		} else {
			fmt.Println("File", fn, "is already in the encryption list.")
		}

		// Encrypt the file
		if !noEncrypt {
			err := encrypter.EncryptFile(fn)
			if err != nil {
				return err
			}
			encryptedCount++
		}

		if !noClean {
			err := encrypter.RemoveFile(fn)
			if err != nil {
				return err
			}
			fmt.Println("Cleaned up plaintext for:", fn)
		}

		err := sourceCodeManager.AddFileToIgnored(fn)
		if err != nil {
			return err
		}
	}

	err := sopsyaml.WriteEncryptFilesToDisk(sopsConfig.Path, sopsConfig.Tree, sopsConfig.EncryptedFiles)
	if err != nil {
		return err
	}

	fmt.Printf("%d file(s) encrypted successfully.\n", encryptedCount)

	return nil
}
