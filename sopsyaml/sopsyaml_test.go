package sopsyaml

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/mozilla-services/yaml"
)

// finds
// not found
// finds a couple levels deep
// hits git and stops
func TestFindConfigFile(t *testing.T) {
	origFs := fs
	t.Run("Finds File Immediate", func(t *testing.T) {
		fs = osFS{
			stat: func(name string) (os.FileInfo, error) {
				return nil, nil
			},
			readfile: func(name string) ([]byte, error) {
				return []byte{}, nil
			},
			writefile: func(name string, data []byte, perms os.FileMode) error {
				return nil
			},
		}
		got, err := FindConfigFile(".")
		if err != nil {
			t.Errorf("FindConfigFile() got err %v", err)
		}
		if got != ".sops.yaml" {
			t.Errorf("FindConfigFile() = %v, want \".sops.yaml\"", got)
		}

		fs = origFs
		return
	})
	t.Run("not found", func(t *testing.T) {
		fs = osFS{
			stat: func(name string) (os.FileInfo, error) {
				return nil, fmt.Errorf("NF")
			},
			readfile: func(name string) ([]byte, error) {
				return []byte{}, nil
			},
			writefile: func(name string, data []byte, perms os.FileMode) error {
				return nil
			},
		}
		_, err := FindConfigFile(".")
		if err == nil {
			t.Errorf("FindConfigFile() expected err")
		}

		fs = origFs
		return
	})
	t.Run("levels deep", func(t *testing.T) {
		fs = osFS{
			stat: func(name string) (os.FileInfo, error) {
				if name == "../../.sops.yaml" {
					return nil, nil
				}
				return nil, fmt.Errorf("NF")
			},
			readfile: func(name string) ([]byte, error) {
				return []byte{}, nil
			},
			writefile: func(name string, data []byte, perms os.FileMode) error {
				return nil
			},
		}
		got, err := FindConfigFile(".")
		if err != nil {
			t.Errorf("FindConfigFile() got err %v", err)
		}
		if got != "../../.sops.yaml" {
			t.Errorf("FindConfigFile() = %v, want \"../../.sops.yaml\"", got)
		}

		fs = origFs
		return
	})
	t.Run("different start", func(t *testing.T) {
		fs = osFS{
			stat: func(name string) (os.FileInfo, error) {
				if name == "directory/.sops.yaml" {
					return nil, nil
				}
				return nil, fmt.Errorf("NF")
			},
			readfile: func(name string) ([]byte, error) {
				return []byte{}, nil
			},
			writefile: func(name string, data []byte, perms os.FileMode) error {
				return nil
			},
		}
		got, err := FindConfigFile("directory/here/goes/further/")
		if err != nil {
			t.Errorf("FindConfigFile() got err %v", err)
		}
		if got != "directory/.sops.yaml" {
			t.Errorf("FindConfigFile() = %v, want \"directory/.sops.yaml\"", got)
		}

		fs = origFs
		return
	})
	t.Run("stops at git repo", func(t *testing.T) {
		fs = osFS{
			stat: func(name string) (os.FileInfo, error) {
				match, err := regexp.MatchString(".git$", name)
				if err != nil {
					return nil, err
				}
				if match {
					return nil, nil
				}
				return nil, fmt.Errorf("NF")
			},
			readfile: func(name string) ([]byte, error) {
				return []byte{}, nil
			},
			writefile: func(name string, data []byte, perms os.FileMode) error {
				return nil
			},
		}
		_, err := FindConfigFile(".")
		if err == nil {
			t.Errorf("FindConfigFile() expected err")
		}

		fs = origFs
		return
	})

}

func TestLoadConfigFile(t *testing.T) {
	origFs := fs
	t.Run("file error", func(t *testing.T) {
		fs = osFS{
			stat: func(name string) (os.FileInfo, error) {
				return nil, nil
			},
			readfile: func(name string) ([]byte, error) {
				return nil, fmt.Errorf("some error")
			},
			writefile: func(name string, data []byte, perms os.FileMode) error {
				return nil
			},
		}

		_, err := LoadConfigFile("filepath")
		if err == nil {
			t.Errorf("LoadConfigFile() expected err")
		}

		fs = origFs
		return
	})
	t.Run("yaml error", func(t *testing.T) {
		fs = osFS{
			stat: func(name string) (os.FileInfo, error) {
				return nil, nil
			},
			readfile: func(name string) ([]byte, error) {
				if name != "filepath" {
					return nil, fmt.Errorf("got %v not filepath", name)
				}
				yml := []byte(`
          ~~~not yaml
          at all
        `)

				return yml, nil
			},
			writefile: func(name string, data []byte, perms os.FileMode) error {
				return nil
			},
		}

		_, err := LoadConfigFile("filepath")
		if err == nil {
			t.Fatalf("LoadConfigFile() expected yaml err")
		}
		if !strings.Contains(err.Error(), "unmarshal") {
			t.Fatalf("expected an unmarshal error")
		}

		fs = origFs
		return
	})
	t.Run("successful parse", func(t *testing.T) {
		fs = osFS{
			stat: func(name string) (os.FileInfo, error) {
				return nil, nil
			},
			readfile: func(filepath string) ([]byte, error) {
				yml := []byte(`
          yaml:
            - in
            - a string
        `)

				return yml, nil
			},
			writefile: func(name string, data []byte, perms os.FileMode) error {
				return nil
			},
		}

		yml, err := LoadConfigFile("filepath")
		if err != nil {
			t.Errorf("LoadConfigFile() got err: %v", err)
		}

		stringRep := fmt.Sprintf("%v", yml)
		if stringRep != "&[{yaml [in a string]}]" {
			t.Errorf("LoadConfigFile() = %v, want \"&[{yaml [in a string]}]\"", stringRep)
		}

		fs = origFs
		return
	})
}

func TestWriteConfigFile(t *testing.T) {
	origFs := fs
	t.Run("cant unmarshal", func(t *testing.T) {
		//TODO what are valid errors here
		t.Skipf("Unsure what would error in yaml yet")

		fs = origFs
		return
	})
	t.Run("cant write", func(t *testing.T) {
		fs = osFS{
			stat: func(name string) (os.FileInfo, error) {
				return nil, nil
			},
			readfile: func(name string) ([]byte, error) {
				return nil, nil
			},
			writefile: func(name string, data []byte, perms os.FileMode) error {
				if name != "filepath" {
					return fmt.Errorf("Expected filepath got %v", name)
				}
				return fmt.Errorf("Error")
			},
		}

		yml := make(yaml.MapSlice, 0)
		err := WriteConfigFile("filepath", &yml)

		if err == nil {
			t.Fatalf("WriteConfigFile() expected an error")
		}
		if err.Error() != "Error" {
			t.Fatalf("WriteConfigFile() unexpected error %v", err)
		}

		fs = origFs
		return
	})
	t.Run("write", func(t *testing.T) {
		fs = osFS{
			stat: func(name string) (os.FileInfo, error) {
				return nil, nil
			},
			readfile: func(name string) ([]byte, error) {
				return nil, nil
			},
			writefile: func(name string, data []byte, perms os.FileMode) error {
				if name != "filepath" {
					return fmt.Errorf("Expected filepath got %v", name)
				}
				if len(data) <= 0 {
					return fmt.Errorf("Got no data")
					//TODO more here?
				}
				if perms != 0644 {
					return fmt.Errorf("Expected 0644 got %v", perms)
				}
				return nil
			},
		}

		yml := yaml.MapSlice{
			yaml.MapItem{Key: "yaml", Value: "one"},
		}

		err := WriteConfigFile("filepath", &yml)
		if err != nil {
			t.Errorf("WriteConfigFile() unexpected error %v", err)
		}

		fs = origFs
		return
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
			name:    "encrypted_files element doesnt exist",
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
	type args struct {
		basePath string
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
	// TODO: Add test cases.
	}

	//todo mock the methods if possible

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConfigEncryptFiles(tt.args.basePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfigEncryptFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfigEncryptFiles() = %v, want %v", got, tt.want)
			}
		})
	}
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
	type args struct {
		confPath string
		data     *yaml.MapSlice
		encFiles []string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}

	//todo mock the methods

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteEncryptFilesToDisk(tt.args.confPath, tt.args.data, tt.args.encFiles); (err != nil) != tt.wantErr {
				t.Errorf("WriteEncryptFilesToDisk() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
