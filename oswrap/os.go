package oswrap

import (
	"os"
	"os/exec"
	"syscall"
)

// OsWrap wraps OS filesystem commands for mocking
type OsWrap interface {
	Command(name string, arg ...string) *exec.Cmd
	Exec(argv0 string, argv []string, envv []string) (err error)
	Create(name string) (*os.File, error)
	LookPath(name string) (string, error)
	Environ() []string
	Remove(name string) error
}

type osWrap struct{}

var ow OsWrap = osWrap{}

// OsWrapInstance gets an instance of the os wrapper
func OsWrapInstance() OsWrap {
	return ow
}

func (ow osWrap) Command(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}
func (ow osWrap) Exec(argv0 string, argv []string, envv []string) (err error) {
	return syscall.Exec(argv0, argv, envv)
}
func (ow osWrap) Create(name string) (*os.File, error) {
	return os.Create(name)
}
func (ow osWrap) LookPath(name string) (string, error) {
	return exec.LookPath(name)
}
func (ow osWrap) Environ() []string {
	return os.Environ()
}
func (ow osWrap) Remove(name string) error {
	return os.Remove(name)
}
