package execwrap

import (
	"testing"

	mock_execwrap "github.com/Ibotta/go-commons/sopstool/execwrap/mock"
	"github.com/golang/mock/gomock"
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
	t.Run("run given command", func(t *testing.T) {
	})
	t.Run("run given with no args", func(t *testing.T) {
	})
	t.Run("run err", func(t *testing.T) {
	})
}

func TestRunCommandStdoutToFile(t *testing.T) {
	t.Run("run given command", func(t *testing.T) {
	})
	t.Run("run given with no args", func(t *testing.T) {
	})
	t.Run("run err", func(t *testing.T) {
	})
	t.Run("file err", func(t *testing.T) {
	})
}

func TestRunSyscallExec(t *testing.T) {
	t.Run("run given command", func(t *testing.T) {
	})
	t.Run("run given with no args", func(t *testing.T) {
	})
}
