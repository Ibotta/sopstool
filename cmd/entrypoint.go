// Copyright © 2017 Ibotta

package cmd

import (
	"fmt"

	"github.com/Ibotta/sopstool/oswrap"
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

	// allowFail is getting inherited from decrypt.go since it is in the same package
	entrypointCmd.Flags().BoolVar(&allowFail, "allow-fail", false, "Do not fail if not all files can be decrypted")
}

// EntrypointCommand the command for the add command
// Note the named return parameter, it is used to tack on errors during
// the deferred encrypted file cleanup.
func EntrypointCommand(cmd *cobra.Command, args []string) (rerr error) {
	initConfig()

	//TODO put this somewhere mockable
	execWrap := oswrap.ExecWrapInstance()

	//TODO don't use the actual commands here.
	err := DecryptCommand(cmd, filesToDecrypt)
	if err != nil {
		return err
	}
	defer func() {
		cleanupErr := CleanCommand(cmd, filesToDecrypt)
		if cleanupErr != nil {
			// Using the named return to stack errors.
			rerr = fmt.Errorf("Encrypted file cleanup error:\n%s\n%s", cleanupErr, rerr)
		}
	}()

	if execCommand {
		execWrap.RunSyscallExec(args)
	} else {
		err := execWrap.RunCommandDirect(args)
		if err != nil {
			return err
		}
	}

	return nil
}
