package sopsyaml

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	mock_oswrap "github.com/Ibotta/sopstool/oswrap/mock"
	"github.com/Ibotta/sopstool/testhelpers"
	"github.com/golang/mock/gomock"
	"github.com/mozilla-services/yaml"
)

func TestFindConfigFile(t *testing.T) {
	origOw := osWrap
	t.Run("Finds File Immediate", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Stat(gomock.Eq(".sops.yaml")).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, nil
		})

		osWrap = mock

		got, err := FindConfigFile(".")
		if err != nil {
			t.Errorf("FindConfigFile() got err %v", err)
		}
		if got != ".sops.yaml" {
			t.Errorf("FindConfigFile() = %v, want \".sops.yaml\"", got)
		}

		osWrap = origOw
	})
	t.Run("not found after never getting Stat", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Stat(testhelpers.RegexMatches(`^.*\.git$`)).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, fmt.Errorf("Not Found") // Never git
		}).AnyTimes()
		mock.EXPECT().Stat(testhelpers.RegexMatches(`^.*\.sops.yaml$`)).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, fmt.Errorf("Not Found")
		}).AnyTimes()

		osWrap = mock

		_, err := FindConfigFile(".")
		if err == nil {
			t.Errorf("FindConfigFile() expected err")
		}

		osWrap = origOw
	})
	t.Run("levels deep", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Stat(testhelpers.RegexMatches(`^.*\.git$`)).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, fmt.Errorf("Not Found") // Never git
		}).AnyTimes()

		mock.EXPECT().Stat(gomock.Eq(".sops.yaml")).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, fmt.Errorf("Not Found")
		})
		mock.EXPECT().Stat(gomock.Eq("../.sops.yaml")).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, fmt.Errorf("Not Found")
		})
		mock.EXPECT().Stat(gomock.Eq("../../.sops.yaml")).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, nil //found
		})

		osWrap = mock

		got, err := FindConfigFile(".")
		if err != nil {
			t.Errorf("FindConfigFile() got err %v", err)
		}
		if got != "../../.sops.yaml" {
			t.Errorf("FindConfigFile() = %v, want \"../../.sops.yaml\"", got)
		}

		osWrap = origOw
	})
	t.Run("different start", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Stat(testhelpers.RegexMatches(`^.*\.git$`)).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, fmt.Errorf("Not Found") // Never git
		}).AnyTimes()

		mock.EXPECT().Stat(gomock.Not(gomock.Eq("directory/.sops.yaml"))).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, fmt.Errorf("Not Found")
		}).AnyTimes()
		mock.EXPECT().Stat(gomock.Eq("directory/.sops.yaml")).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, nil //found
		})

		osWrap = mock

		got, err := FindConfigFile("directory/here/goes/further/")

		if err != nil {
			t.Errorf("FindConfigFile() got err %v", err)
		}
		if got != "directory/.sops.yaml" {
			t.Errorf("FindConfigFile() = %v, want \"directory/.sops.yaml\"", got)
		}

		osWrap = origOw
	})
	t.Run("stops at git repo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Stat(gomock.Eq(".git")).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, nil //find git immediately
		})
		mock.EXPECT().Stat(gomock.Eq(".sops.yaml")).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, fmt.Errorf("Not Found")
		})

		osWrap = mock

		_, err := FindConfigFile(".")
		if err == nil {
			t.Errorf("FindConfigFile() expected err")
		}

		osWrap = origOw
	})
}

func TestLoadConfigFile(t *testing.T) {
	origOw := osWrap
	t.Run("file error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().ReadFile(gomock.Eq("filepath")).DoAndReturn(func(c string) ([]byte, error) {
			return nil, fmt.Errorf("a file read error")
		})

		osWrap = mock

		_, err := LoadConfigFile("filepath")

		if err == nil {
			t.Errorf("LoadConfigFile() expected err")
		}

		osWrap = origOw
	})
	t.Run("yaml error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().ReadFile(gomock.Eq("filepath")).DoAndReturn(func(c string) ([]byte, error) {
			yml := []byte(`
          ~~~not yaml
          at all
        `)

			return yml, nil
		})

		osWrap = mock

		_, err := LoadConfigFile("filepath")

		if err == nil {
			t.Fatalf("LoadConfigFile() expected yaml err")
		}
		if !strings.Contains(err.Error(), "unmarshal") {
			t.Fatalf("expected an unmarshal error")
		}

		osWrap = origOw
	})
	t.Run("successful parse", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().ReadFile(gomock.Eq("filepath")).DoAndReturn(func(c string) ([]byte, error) {
			yml := []byte(`
yaml:
- in
- a string
`)

			return yml, nil
		})

		osWrap = mock

		yml, err := LoadConfigFile("filepath")
		if err != nil {
			t.Errorf("LoadConfigFile() got err: %v", err)
		}

		stringRep := fmt.Sprintf("%v", yml)
		if stringRep != "&[{yaml [in a string]}]" {
			t.Errorf("LoadConfigFile() = %v, want \"&[{yaml [in a string]}]\"", stringRep)
		}

		osWrap = origOw
	})
}

func TestWriteConfigFile(t *testing.T) {
	origOw := osWrap
	t.Run("cant unmarshal", func(t *testing.T) {
		//TODO what are valid errors here
		t.Skipf("Unsure what would error in yaml yet")
	})
	t.Run("cant write", func(t *testing.T) {
		t.Skipf("Reflection of FileMode is messed up")
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)
		expected := []byte(`{}
`)

		// expectedMode := os.FileMode(0644)
		mock.EXPECT().WriteFile(gomock.Eq("filepath"), gomock.Eq(expected), gomock.Eq(0644)).DoAndReturn(func(c string, args ...string) error {
			return fmt.Errorf("a file read error")
		})

		osWrap = mock

		yml := make(yaml.MapSlice, 0)
		err := WriteConfigFile("filepath", &yml)

		if err == nil {
			t.Fatalf("WriteConfigFile() expected an error")
		}
		if err.Error() != "Error" {
			t.Fatalf("WriteConfigFile() unexpected error %v", err)
		}

		osWrap = origOw
	})
	t.Run("write", func(t *testing.T) {
		t.Skipf("Reflection of FileMode is messed up")
		yml := yaml.MapSlice{
			yaml.MapItem{Key: "yaml", Value: "one"},
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)
		expected := []byte(`yaml: one
`)

		mock.EXPECT().WriteFile(gomock.Eq("filepath"), gomock.Eq(expected), gomock.Eq(0644)).DoAndReturn(func(c string, args ...string) error {
			return nil
		})

		osWrap = mock

		err := WriteConfigFile("filepath", &yml)
		if err != nil {
			t.Errorf("WriteConfigFile() unexpected error %v", err)
		}

		osWrap = origOw
	})
}

func unmarshalStringHelper(str string) *yaml.MapSlice {
	var data yaml.MapSlice
	err := (yaml.CommentUnmarshaler{}).Unmarshal([]byte(str), &data)
	if err != nil {
		panic(err)
	}
	return &data
}

func TestExtractConfigEncryptFiles(t *testing.T) {
	type args struct {
		data *yaml.MapSlice
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "encrypted_files not array",
			args:    args{data: unmarshalStringHelper("foo: bar\nencrypted_files: nope")},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "encrypted_files not array of strings",
			args:    args{data: unmarshalStringHelper("foo: bar\nencrypted_files: [1, 2, 3]")},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "encrypted_files element doesn't exist",
			args:    args{data: unmarshalStringHelper("foo: bar")},
			want:    []string{},
			wantErr: false,
		},
		{
			name:    "encrypted_files is empty array",
			args:    args{data: unmarshalStringHelper("foo: bar\nencrypted_files: []")},
			want:    []string{},
			wantErr: false,
		},
		{
			name:    "encrypted_files is array with values",
			args:    args{data: unmarshalStringHelper("foo: bar\nencrypted_files:\n  - first\n  - second")},
			want:    []string{"first", "second"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractConfigEncryptFiles(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractConfigEncryptFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractConfigEncryptFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetConfigEncryptFiles(t *testing.T) {
	origOw := osWrap
	t.Run("load a file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Stat(gomock.Eq(".sops.yaml")).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, nil
		})
		mock.EXPECT().ReadFile(gomock.Eq(".sops.yaml")).DoAndReturn(func(c string) ([]byte, error) {
			yml := []byte(`
foo: bar
encrypted_files:
- one
- two
`)

			return yml, nil
		})

		osWrap = mock

		got, err := GetConfigEncryptFiles(".")

		if err != nil {
			t.Errorf("GetConfigEncryptFiles() got err %v", err)
		}
		expected := []string{"one", "two"}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("GetConfigEncryptFiles() = %v, want %v", got, expected)
		}

		osWrap = origOw
	})
	t.Run("err on file find", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Stat(gomock.Eq(".git")).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, nil //find git immediately
		}).AnyTimes()
		mock.EXPECT().Stat(gomock.Eq(".sops.yaml")).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, fmt.Errorf("Not Found")
		}).AnyTimes()

		osWrap = mock

		_, err := GetConfigEncryptFiles(".")

		if err == nil {
			t.Errorf("GetConfigEncryptFiles() expected an error")
		}

		osWrap = origOw
	})
	t.Run("err on config load", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Stat(gomock.Eq(".sops.yaml")).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, nil
		})
		mock.EXPECT().ReadFile(gomock.Eq(".sops.yaml")).DoAndReturn(func(c string) ([]byte, error) {
			yml := []byte(`~~not good`)

			return yml, nil
		})

		osWrap = mock

		_, err := GetConfigEncryptFiles(".")

		if err == nil {
			t.Errorf("GetConfigEncryptFiles() expected an error")
		}

		osWrap = origOw
	})
	t.Run("err config extract", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		mock.EXPECT().Stat(gomock.Eq(".sops.yaml")).DoAndReturn(func(c string) (*os.FileInfo, error) {
			return nil, nil
		})
		mock.EXPECT().ReadFile(gomock.Eq(".sops.yaml")).DoAndReturn(func(c string) ([]byte, error) {
			yml := []byte(`encrypted_files: [1,2,3]`)

			return yml, nil
		})

		osWrap = mock

		_, err := GetConfigEncryptFiles(".")
		if err == nil {
			t.Errorf("GetConfigEncryptFiles() expected an error")
		}

		osWrap = origOw
	})
}

func TestReplaceConfigEncryptFiles(t *testing.T) {
	type args struct {
		data     *yaml.MapSlice
		encFiles []string
	}

	//use string representation for ease

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Replace existing set",
			args: args{
				data:     unmarshalStringHelper("foo: bar\nencrypted_files:\n  - first\n  - second"),
				encFiles: []string{"one", "two"},
			},
			want:    "&[{foo bar} {encrypted_files [one two]}]",
			wantErr: false,
		},
		{
			name: "Replace with an empty item",
			args: args{
				data:     unmarshalStringHelper("foo: bar\nencrypted_files:\n  - first\n  - second"),
				encFiles: []string{},
			},
			want:    "&[{foo bar} {encrypted_files []}]",
			wantErr: false,
		},
		{
			name: "Replace empty with nonempty",
			args: args{
				data:     unmarshalStringHelper("foo: bar\nencrypted_files: []"),
				encFiles: []string{"one", "two"},
			},
			want:    "&[{foo bar} {encrypted_files [one two]}]",
			wantErr: false,
		},
		{
			name: "adds new top level key",
			args: args{
				data:     unmarshalStringHelper("foo: bar"),
				encFiles: []string{"one", "two"},
			},
			want:    "&[{foo bar} {encrypted_files [one two]}]",
			wantErr: false,
		},
		//todo whats the error case here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReplaceConfigEncryptFiles(tt.args.data, tt.args.encFiles)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReplaceConfigEncryptFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(fmt.Sprintf("%v", got), tt.want) {
				t.Errorf("ReplaceConfigEncryptFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteEncryptFilesToDisk(t *testing.T) {
	origOw := osWrap
	t.Run("write a file", func(t *testing.T) {
		t.Skipf("Reflection of FileMode is messed up")
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		expected := []byte("foo: bar\nencrypted_files:\n- first\n- second\n")
		mock.EXPECT().WriteFile(gomock.Eq("filepath"), gomock.Eq(expected), gomock.Eq(0644)).DoAndReturn(func(c string, args ...string) error {
			return nil
		})

		osWrap = mock

		data := unmarshalStringHelper("foo: bar")
		err := WriteEncryptFilesToDisk("filepath", data, []string{"first", "second"})
		if err != nil {
			t.Errorf("GetConfigEncryptFiles() got err %v", err)
		}

		osWrap = origOw
	})
	t.Run("file write error", func(t *testing.T) {
		t.Skipf("Reflection of FileMode is messed up")
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mock := mock_oswrap.NewMockOsWrap(ctrl)

		expected := []byte("foo: bar\nencrypted_files:\n- first\n- second\n")
		mock.EXPECT().WriteFile(gomock.Eq("filepath"), gomock.Eq(expected), gomock.Eq(0644)).DoAndReturn(func(c string, args ...string) error {
			return fmt.Errorf("some write error")
		})

		osWrap = mock

		data := unmarshalStringHelper("foo: bar")
		err := WriteEncryptFilesToDisk("filepath", data, []string{"first", "second"})
		if err == nil {
			t.Fatalf("GetConfigEncryptFiles() expected an error")
		}

		osWrap = origOw
	})
}
