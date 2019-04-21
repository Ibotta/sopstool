package oswrap

import (
	"errors"
	"os"
	"os/exec"
	"testing"

	mock_oswrap "github.com/Ibotta/sopstool/oswrap/mock"

	"github.com/golang/mock/gomock"
)

func TestRunCommandDirect(t *testing.T) {
	origOW := ow
	t.Run("run given command", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Command(gomock.Eq("sops"), gomock.Eq("myfile.sops.yaml")).DoAndReturn(func(c string, args ...string) *exec.Cmd {
			return exec.Command("true")
		})

		ow = mock

		err := ew.RunCommandDirect([]string{"sops", "myfile.sops.yaml"})

		if err != nil {
			t.Errorf("TestRunCommandDirect() unexpected error %v", err)
		}

		ow = origOW
		return
	})
	t.Run("run given with no args", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Command(gomock.Eq("sops")).DoAndReturn(func(c string, args ...string) *exec.Cmd {
			return exec.Command("true")
		})

		ow = mock

		err := ew.RunCommandDirect([]string{"sops"})

		if err != nil {
			t.Errorf("TestRunCommandDirect() unexpected error %v", err)
		}

		ow = origOW
		return
	})
	t.Run("run err", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Command(gomock.Eq("sops"), gomock.Eq("myfile.sops.yaml")).DoAndReturn(func(c string, args ...string) *exec.Cmd {
			return exec.Command("false")
		})

		ow = mock
		err := ew.RunCommandDirect([]string{"sops", "myfile.sops.yaml"})

		if err == nil {
			t.Errorf("TestRunCommandDirect() expected error, got %v", err)
		}

		ow = origOW
		return
	})
}

func TestRunCommandStdoutToFile(t *testing.T) {
	origOW := ow
	t.Run("run given command", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Command(gomock.Eq("sops"), gomock.Eq("myfile.sops.yaml")).DoAndReturn(func(c string, args ...string) *exec.Cmd {
			return exec.Command("true")
		})

		mock.EXPECT().Create(gomock.Eq("filename")).DoAndReturn(func(f string) (*os.File, error) {
			//TODO replace all file stuff with afero
			return os.Create("/tmp/TestRunCommandStdoutToFile")
		})
		defer os.Remove("/tmp/TestRunCommandStdoutToFile")

		ow = mock

		err := ew.RunCommandStdoutToFile("filename", []string{"sops", "myfile.sops.yaml"})

		if err != nil {
			t.Errorf("TestRunCommandStdoutToFile() unexpected error %v", err)
		}

		ow = origOW
		return
	})
	t.Run("run given with no args", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Command(gomock.Eq("sops")).DoAndReturn(func(c string, args ...string) *exec.Cmd {
			return exec.Command("true")
		})

		mock.EXPECT().Create(gomock.Eq("filename")).DoAndReturn(func(f string) (*os.File, error) {
			//TODO replace all file stuff with afero
			return os.Create("/tmp/TestRunCommandStdoutToFile")
		})
		defer os.Remove("/tmp/TestRunCommandStdoutToFile")

		ow = mock

		err := ew.RunCommandStdoutToFile("filename", []string{"sops"})

		if err != nil {
			t.Errorf("TestRunCommandStdoutToFile() unexpected error %v", err)
		}

		ow = origOW
		return
	})
	t.Run("run err", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Command(gomock.Eq("sops"), gomock.Eq("myfile.sops.yaml")).DoAndReturn(func(c string, args ...string) *exec.Cmd {
			return exec.Command("false")
		})

		mock.EXPECT().Create(gomock.Eq("filename")).DoAndReturn(func(f string) (*os.File, error) {
			//TODO replace all file stuff with afero
			return os.Create("/tmp/TestRunCommandStdoutToFile")
		})
		defer os.Remove("/tmp/TestRunCommandStdoutToFile")

		mock.EXPECT().Remove(gomock.Eq("filename")).Return(nil)

		ow = mock

		err := ew.RunCommandStdoutToFile("filename", []string{"sops", "myfile.sops.yaml"})

		if err == nil {
			t.Errorf("TestRunCommandStdoutToFile() expected error, got %v", err)
		}

		ow = origOW
		return
	})
	t.Run("file err", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Command(gomock.Eq("sops"), gomock.Eq("myfile.sops.yaml")).DoAndReturn(func(c string, args ...string) *exec.Cmd {
			return exec.Command("false")
		})

		mock.EXPECT().Create(gomock.Eq("filename")).Return(nil, errors.New("an error"))

		ow = mock

		err := ew.RunCommandStdoutToFile("filename", []string{"sops", "myfile.sops.yaml"})

		if err == nil {
			t.Errorf("TestRunCommandStdoutToFile() expected error, got %v", err)
		}

		ow = origOW
		return
	})
	// TODO test when Close err, but requires mocking of file
}

func TestRunSyscallExec(t *testing.T) {
	origOW := ow
	t.Run("run given command", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().LookPath(gomock.Eq("sops")).Return("sopspath", nil)
		mock.EXPECT().Environ().Return([]string{"one"})
		mock.EXPECT().Exec(gomock.Eq("sopspath"), gomock.Eq([]string{"sops", "myfile.sops.yaml"}), []string{"one"}).Return(nil)

		ow = mock

		err := ew.RunSyscallExec([]string{"sops", "myfile.sops.yaml"})

		if err != nil {
			t.Errorf("TestRunSyscallExec() unexpected error %v", err)
		}

		ow = origOW
		return
	})
	t.Run("run given with no args", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().LookPath(gomock.Eq("sops")).Return("sopspath", nil)
		mock.EXPECT().Environ().Return([]string{"one"})
		mock.EXPECT().Exec(gomock.Eq("sopspath"), gomock.Eq([]string{"sops"}), []string{"one"}).Return(nil)

		ow = mock

		err := ew.RunSyscallExec([]string{"sops"})

		if err != nil {
			t.Errorf("TestRunSyscallExec() unexpected error %v", err)
		}

		ow = origOW
		return
	})
	t.Run("exec error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().LookPath(gomock.Eq("sops")).Return("sopspath", nil)
		mock.EXPECT().Environ().Return([]string{"one"})
		mock.EXPECT().Exec(gomock.Eq("sopspath"), gomock.Eq([]string{"sops", "myfile.sops.yaml"}), []string{"one"}).Return(errors.New("an error"))

		ow = mock

		err := ew.RunSyscallExec([]string{"sops", "myfile.sops.yaml"})

		if err == nil {
			t.Errorf("TestRunSyscallExec() expected error, got %v", err)
		}

		ow = origOW
		return
	})
	t.Run("lookpath error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().LookPath(gomock.Eq("sops")).Return("", errors.New("an error"))
		ow = mock

		err := ew.RunSyscallExec([]string{"sops", "myfile.sops.yaml"})

		if err == nil {
			t.Errorf("TestRunSyscallExec() expected error, got %v", err)
		}

		ow = origOW
		return
	})
}
