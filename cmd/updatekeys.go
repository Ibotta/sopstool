// Copyright © 2017 Ibotta

package cmd

import (
	"github.com/Ibotta/sopstool/fileutil"
	"github.com/spf13/cobra"
)

// updatekeysCmd represents the updatekeysCmd command
var updatekeysCmd = &cobra.Command{
	Use:   "updatekeys [files ...]",
	Short: "update recipients keys",
	Long:  "update recipients keys",
	RunE:  UpdateKeysCommand,
}

var nonInteractiveMode bool

func init() {
	RootCmd.AddCommand(updatekeysCmd)
	updatekeysCmd.Flags().BoolVar(&nonInteractiveMode, "non-interactive", false, "pre-approve all changes and run non-interactively")
}

// UpdateKeysCommand updates recipients keys
func UpdateKeysCommand(cmd *cobra.Command, args []string) error {
	initConfig()
	var extraArgs []string

	filesToUpdate, err := fileutil.SomeOrAllFiles(args, sopsConfig.EncryptedFiles)
	if err != nil {
		return err
	}

	if nonInteractiveMode {
		extraArgs = append(extraArgs, "--yes")
	}

	//Update recipients keys on files
	for _, f := range filesToUpdate {
		err := encrypter.UpdateKeysFile(f, extraArgs)
		if err != nil {
			return err
		}
	}

	return nil
}
