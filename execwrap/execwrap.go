package execwrap

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/Ibotta/go-commons/sopstool/fileutil"
)

// EncryptFile encrypts a file rewriting the sops encrypted file
func EncryptFile(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)
	return RunCommandStdoutToFile(cryptfile, []string{"sops", "-e", fn})
}

// DecryptFile decrypts a file rewriting the plaintext file
func DecryptFile(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)
	return RunCommandStdoutToFile(fn, []string{"sops", "-d", cryptfile})
}

// DecryptFilePrint decrypts a file printing the result
func DecryptFilePrint(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)
	return RunCommandDirect([]string{"sops", "-d", cryptfile})
}

// RemoveFile removes a plaintext file from the filesystem
func RemoveFile(fn string) error {
	return RunCommandDirect([]string{"rm", fn})
}

// RemoveSopsFile removes a sops file from the filesystem
func RemoveSopsFile(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)

	return RunCommandDirect([]string{"rm", cryptfile})
}

// RotateFile rotates keys on a file
func RotateFile(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)

	return RunCommandDirect([]string{"sops", "-i", "-r", cryptfile})
}

// EditFile should open the editor for a file
func EditFile(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)

	return RunCommandDirect([]string{"sops", cryptfile})
}

// RunCommandDirect runs a command, redirecting 0/1/2 to the caller
func RunCommandDirect(command []string) error {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Start()
	if err != nil {
		return err
	}
	cmd.Wait()

	return nil
}

// RunCommandStdoutToFile runs a command, redirecting Stdout to a file, the rest to caller
func RunCommandStdoutToFile(outfileName string, command []string) error {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	// open the out file for writing
	outfile, err := os.Create(outfileName)
	if err != nil {
		return err
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	err = cmd.Start()
	if err != nil {
		return err
	}
	cmd.Wait()

	return nil
}

// RunSyscallExec runs exec which fully takes over the process
func RunSyscallExec(args []string) error {
	path, err := exec.LookPath(args[0])
	if err != nil {
		return err
	}

	return syscall.Exec(path, args, os.Environ())
}
