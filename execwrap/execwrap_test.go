package execwrap

import (
	"errors"
	"os"
	"os/exec"
	"testing"

	mock_execwrap "github.com/Ibotta/sopstool/execwrap/mock"
	"github.com/golang/mock/gomock"
	// "github.com/spf13/afero"
)

type mockExecWrap struct{}

func TestEncryptFile(t *testing.T) {
	origEw := ew
	t.Run("run enc", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMockexecutionWrapper(ctrl)

		mock.EXPECT().RunCommandStdoutToFile("myfile.sops.yaml", gomock.Eq([]string{"sops", "-e", "myfile.yaml"})).Return(nil)

		ew = mock

		err := EncryptFile("myfile.yaml")
		if err != nil {
			t.Errorf("TestEncryptFile() unexpected error %v", err)
		}

		ew = origEw
		return
	})
}

func TestDecryptFile(t *testing.T) {
	origEw := ew
	t.Run("run dec", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMockexecutionWrapper(ctrl)

		mock.EXPECT().RunCommandStdoutToFile("myfile.yaml", gomock.Eq([]string{"sops", "-d", "myfile.sops.yaml"})).Return(nil)

		ew = mock

		err := DecryptFile("myfile.yaml")
		if err != nil {
			t.Errorf("TestEncryptFile() unexpected error %v", err)
		}

		ew = origEw
		return
	})
	t.Run("run dec returns error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMockexecutionWrapper(ctrl)

		mock.EXPECT().RunCommandStdoutToFile("myfile.yaml", gomock.Eq([]string{"sops", "-d", "myfile.sops.yaml"})).Return(errors.New("did someting bad"))

		ew = mock

		err := DecryptFile("myfile.yaml")
		if err == nil {
			t.Errorf("TestEncryptFile() expected an error, got %v", err)
		}

		ew = origEw
		return
	})
}

func TestDecryptFilePrint(t *testing.T) {
	origEw := ew
	t.Run("run dec print", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMockexecutionWrapper(ctrl)

		mock.EXPECT().RunCommandDirect(gomock.Eq([]string{"sops", "-d", "myfile.sops.yaml"})).Return(nil)

		ew = mock

		err := DecryptFilePrint("myfile.yaml")
		if err != nil {
			t.Errorf("TestEncryptFile() unexpected error %v", err)
		}

		ew = origEw
		return
	})
}

func TestRemoveFile(t *testing.T) {
	origEw := ew
	t.Run("run rm", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMockexecutionWrapper(ctrl)

		mock.EXPECT().RunCommandDirect(gomock.Eq([]string{"rm", "myfile.yaml"})).Return(nil)

		ew = mock

		err := RemoveFile("myfile.yaml")
		if err != nil {
			t.Errorf("TestEncryptFile() unexpected error %v", err)
		}

		ew = origEw
		return
	})
}

func TestRemoveSopsFile(t *testing.T) {
	origEw := ew
	t.Run("run rm", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMockexecutionWrapper(ctrl)

		mock.EXPECT().RunCommandDirect(gomock.Eq([]string{"rm", "myfile.sops.yaml"})).Return(nil)

		ew = mock

		err := RemoveSopsFile("myfile.yaml")
		if err != nil {
			t.Errorf("TestEncryptFile() unexpected error %v", err)
		}

		ew = origEw
		return
	})
}

func TestRotateFile(t *testing.T) {
	origEw := ew
	t.Run("run rotate", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMockexecutionWrapper(ctrl)

		mock.EXPECT().RunCommandDirect(gomock.Eq([]string{"sops", "-i", "-r", "myfile.sops.yaml"})).Return(nil)
		ew = mock

		err := RotateFile("myfile.yaml")
		if err != nil {
			t.Errorf("TestEncryptFile() unexpected error %v", err)
		}

		ew = origEw
		return
	})
}

func TestEditFile(t *testing.T) {
	origEw := ew
	t.Run("run edit", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMockexecutionWrapper(ctrl)

		mock.EXPECT().RunCommandDirect(gomock.Eq([]string{"sops", "myfile.sops.yaml"})).Return(nil)
		ew = mock

		err := EditFile("myfile.yaml")
		if err != nil {
			t.Errorf("TestEncryptFile() unexpected error %v", err)
		}

		ew = origEw
		return
	})
}

func TestRunCommandDirect(t *testing.T) {
	origE := e
	t.Run("run given command", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMocksystemExec(ctrl)

		mock.EXPECT().Command(gomock.Eq("sops"), gomock.Eq("myfile.sops.yaml")).DoAndReturn(func(c string, args ...string) *exec.Cmd {
			return exec.Command("true")
		})

		e = mock

		err := ew.RunCommandDirect([]string{"sops", "myfile.sops.yaml"})

		if err != nil {
			t.Errorf("TestRunCommandDirect() unexpected error %v", err)
		}

		e = origE
		return
	})
	t.Run("run given with no args", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMocksystemExec(ctrl)

		mock.EXPECT().Command(gomock.Eq("sops")).DoAndReturn(func(c string, args ...string) *exec.Cmd {
			return exec.Command("true")
		})

		e = mock

		err := ew.RunCommandDirect([]string{"sops"})

		if err != nil {
			t.Errorf("TestRunCommandDirect() unexpected error %v", err)
		}

		e = origE
		return
	})
	t.Run("run err", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMocksystemExec(ctrl)

		mock.EXPECT().Command(gomock.Eq("sops"), gomock.Eq("myfile.sops.yaml")).DoAndReturn(func(c string, args ...string) *exec.Cmd {
			return exec.Command("false")
		})

		e = mock

		err := ew.RunCommandDirect([]string{"sops", "myfile.sops.yaml"})

		if err == nil {
			t.Errorf("TestRunCommandDirect() expected error, got %v", err)
		}

		e = origE
		return
	})
}

func TestRunCommandStdoutToFile(t *testing.T) {
	origE := e
	t.Run("run given command", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMocksystemExec(ctrl)

		mock.EXPECT().Command(gomock.Eq("sops"), gomock.Eq("myfile.sops.yaml")).DoAndReturn(func(c string, args ...string) *exec.Cmd {
			return exec.Command("true")
		})

		mock.EXPECT().Create(gomock.Eq("filename")).DoAndReturn(func(f string) (*os.File, error) {
			//TODO replace all file stuff with afero
			return os.Create("/tmp/TestRunCommandStdoutToFile")
		})
		defer os.Remove("/tmp/TestRunCommandStdoutToFile")

		e = mock

		err := ew.RunCommandStdoutToFile("filename", []string{"sops", "myfile.sops.yaml"})

		if err != nil {
			t.Errorf("TestRunCommandStdoutToFile() unexpected error %v", err)
		}

		e = origE
		return
	})
	t.Run("run given with no args", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMocksystemExec(ctrl)

		mock.EXPECT().Command(gomock.Eq("sops")).DoAndReturn(func(c string, args ...string) *exec.Cmd {
			return exec.Command("true")
		})

		mock.EXPECT().Create(gomock.Eq("filename")).DoAndReturn(func(f string) (*os.File, error) {
			//TODO replace all file stuff with afero
			return os.Create("/tmp/TestRunCommandStdoutToFile")
		})
		defer os.Remove("/tmp/TestRunCommandStdoutToFile")

		e = mock

		err := ew.RunCommandStdoutToFile("filename", []string{"sops"})

		if err != nil {
			t.Errorf("TestRunCommandStdoutToFile() unexpected error %v", err)
		}

		e = origE
		return
	})
	t.Run("run err", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMocksystemExec(ctrl)

		mock.EXPECT().Command(gomock.Eq("sops"), gomock.Eq("myfile.sops.yaml")).DoAndReturn(func(c string, args ...string) *exec.Cmd {
			return exec.Command("false")
		})

		mock.EXPECT().Create(gomock.Eq("filename")).DoAndReturn(func(f string) (*os.File, error) {
			//TODO replace all file stuff with afero
			return os.Create("/tmp/TestRunCommandStdoutToFile")
		})
		defer os.Remove("/tmp/TestRunCommandStdoutToFile")

		e = mock

		err := ew.RunCommandStdoutToFile("filename", []string{"sops", "myfile.sops.yaml"})

		if err == nil {
			t.Errorf("TestRunCommandStdoutToFile() expected error, got %v", err)
		}

		e = origE
		return
	})
	t.Run("file err", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMocksystemExec(ctrl)

		mock.EXPECT().Command(gomock.Eq("sops"), gomock.Eq("myfile.sops.yaml")).DoAndReturn(func(c string, args ...string) *exec.Cmd {
			return exec.Command("false")
		})

		mock.EXPECT().Create(gomock.Eq("filename")).Return(nil, errors.New("an error"))

		e = mock

		err := ew.RunCommandStdoutToFile("filename", []string{"sops", "myfile.sops.yaml"})

		if err == nil {
			t.Errorf("TestRunCommandStdoutToFile() expected error, got %v", err)
		}

		e = origE
		return
	})
}

func TestRunSyscallExec(t *testing.T) {
	origE := e
	t.Run("run given command", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMocksystemExec(ctrl)

		mock.EXPECT().LookPath(gomock.Eq("sops")).Return("sopspath", nil)
		mock.EXPECT().Environ().Return([]string{"one"})
		mock.EXPECT().Exec(gomock.Eq("sopspath"), gomock.Eq([]string{"sops", "myfile.sops.yaml"}), []string{"one"}).Return(nil)

		e = mock

		err := ew.RunSyscallExec([]string{"sops", "myfile.sops.yaml"})

		if err != nil {
			t.Errorf("TestRunSyscallExec() unexpected error %v", err)
		}

		e = origE
		return
	})
	t.Run("run given with no args", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMocksystemExec(ctrl)

		mock.EXPECT().LookPath(gomock.Eq("sops")).Return("sopspath", nil)
		mock.EXPECT().Environ().Return([]string{"one"})
		mock.EXPECT().Exec(gomock.Eq("sopspath"), gomock.Eq([]string{"sops"}), []string{"one"}).Return(nil)

		e = mock

		err := ew.RunSyscallExec([]string{"sops"})

		if err != nil {
			t.Errorf("TestRunSyscallExec() unexpected error %v", err)
		}

		e = origE
		return
	})
	t.Run("exec error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMocksystemExec(ctrl)

		mock.EXPECT().LookPath(gomock.Eq("sops")).Return("sopspath", nil)
		mock.EXPECT().Environ().Return([]string{"one"})
		mock.EXPECT().Exec(gomock.Eq("sopspath"), gomock.Eq([]string{"sops", "myfile.sops.yaml"}), []string{"one"}).Return(errors.New("an error"))

		e = mock

		err := ew.RunSyscallExec([]string{"sops", "myfile.sops.yaml"})

		if err == nil {
			t.Errorf("TestRunSyscallExec() expected error, got %v", err)
		}

		e = origE
		return
	})
	t.Run("lookpath error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_execwrap.NewMocksystemExec(ctrl)

		mock.EXPECT().LookPath(gomock.Eq("sops")).Return("", errors.New("an error"))
		e = mock

		err := ew.RunSyscallExec([]string{"sops", "myfile.sops.yaml"})

		if err == nil {
			t.Errorf("TestRunSyscallExec() expected error, got %v", err)
		}

		e = origE
		return
	})
}
