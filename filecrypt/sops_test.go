package filecrypt

import (
	"errors"
	"testing"

	mock_oswrap "github.com/Ibotta/sopstool/oswrap/mock"
	"github.com/golang/mock/gomock"
)

func TestEncryptFile(t *testing.T) {
	origEw := sops.execWrap
	t.Run("run enc", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockExecWrap(ctrl)

		mock.EXPECT().RunCommandStdoutToFile("myfile.sops.yaml", gomock.Eq([]string{"sops", "-e", "myfile.yaml"})).Return(nil)

		sops.execWrap = mock

		err := sops.EncryptFile("myfile.yaml")
		if err != nil {
			t.Errorf("TestEncryptFile() unexpected error %v", err)
		}

		sops.execWrap = origEw
	})
}

func TestDecryptFile(t *testing.T) {
	origEw := sops.execWrap
	t.Run("run dec", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockExecWrap(ctrl)

		mock.EXPECT().RunCommandStdoutToFile("myfile.yaml", gomock.Eq([]string{"sops", "-d", "myfile.sops.yaml"})).Return(nil)

		sops.execWrap = mock

		err := sops.DecryptFile("myfile.yaml")
		if err != nil {
			t.Errorf("TestEncryptFile() unexpected error %v", err)
		}

		sops.execWrap = origEw
	})
	t.Run("run dec returns error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockExecWrap(ctrl)

		mock.EXPECT().RunCommandStdoutToFile("myfile.yaml", gomock.Eq([]string{"sops", "-d", "myfile.sops.yaml"})).Return(errors.New("did someting bad"))

		sops.execWrap = mock

		err := sops.DecryptFile("myfile.yaml")
		if err == nil {
			t.Errorf("TestEncryptFile() expected an error, got %v", err)
		}

		sops.execWrap = origEw
	})
}

func TestDecryptFilePrint(t *testing.T) {
	origEw := sops.execWrap
	t.Run("run dec print", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockExecWrap(ctrl)

		mock.EXPECT().RunCommandDirect(gomock.Eq([]string{"sops", "-d", "myfile.sops.yaml"})).Return(nil)

		sops.execWrap = mock

		err := sops.DecryptFilePrint("myfile.yaml")
		if err != nil {
			t.Errorf("TestEncryptFile() unexpected error %v", err)
		}

		sops.execWrap = origEw
	})
}

func TestRemoveFile(t *testing.T) {
	origOw := sops.osWrap
	t.Run("removes", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Remove(gomock.Eq("myfile.yaml")).Return(nil)

		sops.osWrap = mock

		err := sops.RemoveFile("myfile.yaml")
		if err != nil {
			t.Errorf("TestEncryptFile() unexpected error %v", err)
		}

		sops.osWrap = origOw
	})
}

func TestRemoveCryptFile(t *testing.T) {
	origOw := sops.osWrap
	t.Run("run rm", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Remove(gomock.Eq("myfile.sops.yaml")).Return(nil)

		sops.osWrap = mock

		err := sops.RemoveCryptFile("myfile.yaml")
		if err != nil {
			t.Errorf("TestEncryptFile() unexpected error %v", err)
		}

		sops.osWrap = origOw
	})
}

func TestRotateFile(t *testing.T) {
	origEw := sops.execWrap
	t.Run("run rotate", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockExecWrap(ctrl)

		mock.EXPECT().RunCommandDirect(gomock.Eq([]string{"sops", "-i", "-r", "myfile.sops.yaml"})).Return(nil)
		sops.execWrap = mock

		err := sops.RotateFile("myfile.yaml")
		if err != nil {
			t.Errorf("TestEncryptFile() unexpected error %v", err)
		}

		sops.execWrap = origEw
	})
}

func TestEditFile(t *testing.T) {
	origEw := sops.execWrap
	t.Run("run edit", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockExecWrap(ctrl)

		mock.EXPECT().RunCommandDirect(gomock.Eq([]string{"sops", "myfile.sops.yaml"})).Return(nil)
		sops.execWrap = mock

		err := sops.EditFile("myfile.yaml")
		if err != nil {
			t.Errorf("TestEncryptFile() unexpected error %v", err)
		}

		sops.execWrap = origEw
	})
}
