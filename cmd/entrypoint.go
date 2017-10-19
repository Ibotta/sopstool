// Copyright © 2017 Ibotta

package cmd

import (
	"github.com/Ibotta/go-commons/sopstool/execwrap"
	"github.com/spf13/cobra"
)

// entrypointCmd represents the entrypoint command
var entrypointCmd = &cobra.Command{
	Aliases: []string{"e", "enter"},
	Use:     "entrypoint",
	Short:   "execute a command with decrypted files",
	Long:    `Decrypt files before run, run, cleanup`,
	Args:    cobra.MinimumNArgs(1),
	RunE:    EntrypointCommand,
}

var execCommand bool
var filesToDecrypt []string

func init() {
	RootCmd.AddCommand(entrypointCmd)
	entrypointCmd.Flags().BoolVarP(&execCommand, "exec", "e", false, "Delegate to the command directly with exec(3), no cleanup")
	entrypointCmd.Flags().StringSliceVarP(&filesToDecrypt, "files", "f", []string{}, "files to decrypt (default all)")
}

// EntrypointCommand the command for the add command
func EntrypointCommand(cmd *cobra.Command, args []string) error {
	err := DecryptCommand(cmd, filesToDecrypt)
	if err != nil {
		return err
	}

	if execCommand {
		execwrap.RunSyscallExec(args)
	} else {
		err := execwrap.RunCommandDirect(args)
		if err != nil {
			return err
		}

		err = CleanCommand(cmd, filesToDecrypt)
		if err != nil {
			return err
		}
	}

	return nil
}
