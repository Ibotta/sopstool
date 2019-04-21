package oswrap

import (
	"os"
)

// ExecWrap wraps exec calls
type ExecWrap interface {
	RunCommandDirect(command []string) error
	RunCommandStdoutToFile(outfileName string, command []string) error
	RunSyscallExec(args []string) error
}

type execWrap struct{}

//todo use OsWrapInstance() instead of package local ow?
var ew ExecWrap = execWrap{}

// ExecWrapInstance gets the execution wrapper interface
func ExecWrapInstance() ExecWrap {
	return ew
}

// RunCommandDirect runs a command, redirecting 0/1/2 to the caller
func (ew execWrap) RunCommandDirect(command []string) error {
	cmd := ow.Command(command[0], command[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Wait()
}

// RunCommandStdoutToFile runs a command, redirecting Stdout to a file, the rest to caller
func (ew execWrap) RunCommandStdoutToFile(outfileName string, command []string) error {
	cmd := ow.Command(command[0], command[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	// open the out file for writing
	outfile, err := ow.Create(outfileName)
	if err != nil {
		return err
	}
	cmd.Stdout = outfile

	err = cmd.Start()
	if err != nil {
		// todo defer?
		outfile.Close()
		// todo defer?
		ow.Remove(outfileName)
		return err
	}

	ret := cmd.Wait()
	err = outfile.Close()
	if err != nil {
		return err
	}
	if ret != nil {
		// todo defer?
		ow.Remove(outfileName)
	}

	return ret
}

// RunSyscallExec runs exec which fully takes over the process.
// the return here never really fires unless the exec fails
func (ew execWrap) RunSyscallExec(args []string) error {
	path, err := ow.LookPath(args[0])
	if err != nil {
		return err
	}

	return ow.Exec(path, args, ow.Environ())
}
