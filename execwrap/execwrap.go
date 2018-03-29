package execwrap

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/Ibotta/sopstool/fileutil"
)

// wrap OS filesystem commands for mocking
type systemExec interface {
	Command(name string, arg ...string) *exec.Cmd
	Exec(argv0 string, argv []string, envv []string) (err error)
	Create(name string) (*os.File, error)
	LookPath(name string) (string, error)
	Environ() []string
	Remove(name string) error
}
type osExec struct {
	command  func(name string, arg ...string) *exec.Cmd
	exec     func(argv0 string, argv []string, envv []string) (err error)
	create   func(name string) (*os.File, error)
	lookPath func(name string) (string, error)
	environ  func() []string
	remove   func(name string) error
}

// TODO simplify this now that we use gomock
func (e osExec) Command(name string, arg ...string) *exec.Cmd {
	return e.command(name, arg...)
}
func (e osExec) Exec(argv0 string, argv []string, envv []string) (err error) {
	return e.exec(argv0, argv, envv)
}
func (e osExec) Create(name string) (*os.File, error) {
	return e.create(name)
}
func (e osExec) LookPath(name string) (string, error) {
	return e.lookPath(name)
}
func (e osExec) Environ() []string {
	return e.environ()
}
func (e osExec) Remove(name string) error {
	return e.remove(name)
}

var e systemExec = osExec{
	command:  exec.Command,
	exec:     syscall.Exec,
	create:   os.Create,
	lookPath: exec.LookPath,
	environ:  os.Environ,
	remove:   os.Remove,
}

// EncryptFile encrypts a file rewriting the sops encrypted file
func EncryptFile(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)
	return ew.RunCommandStdoutToFile(cryptfile, []string{"sops", "-e", fn})
}

// DecryptFile decrypts a file rewriting the plaintext file
func DecryptFile(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)
	return ew.RunCommandStdoutToFile(fn, []string{"sops", "-d", cryptfile})
}

// DecryptFilePrint decrypts a file printing the result
func DecryptFilePrint(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)
	return ew.RunCommandDirect([]string{"sops", "-d", cryptfile})
}

// RemoveFile removes a plaintext file from the filesystem
func RemoveFile(fn string) error {
	return e.Remove(fn)
}

// RemoveSopsFile removes a sops file from the filesystem
func RemoveSopsFile(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)

	return e.Remove(cryptfile)
}

// RotateFile rotates keys on a file
func RotateFile(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)

	return ew.RunCommandDirect([]string{"sops", "-i", "-r", cryptfile})
}

// EditFile should open the editor for a file
func EditFile(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)

	return ew.RunCommandDirect([]string{"sops", cryptfile})
}

//wrap the more complex execution wrappers so they're simple to mock

//ExecutionWrapper wraps exec calls
type ExecutionWrapper interface {
	RunCommandDirect(command []string) error
	RunCommandStdoutToFile(outfileName string, command []string) error
	RunSyscallExec(args []string) error
}

type execWrap struct{}

var ew ExecutionWrapper = execWrap{}

// ExecWrap gets the execution wrapper interface
func ExecWrap() ExecutionWrapper {
	return ew
}

// RunCommandDirect runs a command, redirecting 0/1/2 to the caller
func (ew execWrap) RunCommandDirect(command []string) error {
	cmd := e.Command(command[0], command[1:]...)
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
	cmd := e.Command(command[0], command[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	// open the out file for writing
	outfile, err := e.Create(outfileName)
	if err != nil {
		return err
	}
	cmd.Stdout = outfile

	err = cmd.Start()
	if err != nil {
		outfile.Close()
		e.Remove(outfileName)
		return err
	}

	ret := cmd.Wait()
	err = outfile.Close()
	if err != nil {
		return err
	}
	if ret != nil {
		e.Remove(outfileName)
	}

	return ret
}

// RunSyscallExec runs exec which fully takes over the process.
// the return here never really fires unless the exec fails
func (ew execWrap) RunSyscallExec(args []string) error {
	path, err := e.LookPath(args[0])
	if err != nil {
		return err
	}

	return e.Exec(path, args, e.Environ())
}
