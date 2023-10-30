// Copyright © 2017 Ibotta

package cmd

import (
	"fmt"
	"os"

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

	decryptedCount := 0 // Counter for successfully decrypted files

	// Decrypt all the files
	for _, f := range filesToDecrypt {
		cryptfile := fileutil.NormalizeToSopsFile(f)
		if _, err := os.Stat(cryptfile); os.IsNotExist(err) {
			fmt.Println("Encrypted version of file does not exist:", f)
			// If the encrypted file doesn't exist, skip decryption for this file
			continue
		}

		err := encrypter.DecryptFile(f)
		if err != nil {
			if !allowFail {
				return err
			}
		} else {
			decryptedCount++ // Increment the counter when decryption is successful
		}
	}

	// Print out the total number of decrypted files
	fmt.Printf("%d files decrypted successfully.\n", decryptedCount)

	return nil
}
