// Copyright © 2017 Ibotta

package cmd

import (
	"github.com/Ibotta/go-commons/sopstool/execwrap"
	"github.com/Ibotta/go-commons/sopstool/fileutil"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean [files ...]",
	Short: "cleanup plaintext files",
	Long:  `Cleanup the plaintext of some or all files`,
	RunE:  CleanCommand,
}

func init() {
	RootCmd.AddCommand(cleanCmd)
}

// CleanCommand cleans up files
func CleanCommand(cmd *cobra.Command, args []string) error {
	filesToClean, err := fileutil.SomeOrAllFiles(args, sopsConfig.EncryptedFiles)
	if err != nil {
		return err
	}

	//clean all the files
	for _, f := range filesToClean {
		err := execwrap.RemoveFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
