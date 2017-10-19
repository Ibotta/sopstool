package sopsyaml_test

import (
	"reflect"
	"testing"

	"github.com/Ibotta/go-commons/sopstool/sopsyaml"
	"github.com/mozilla-services/yaml"
)

func TestFindConfigFile(t *testing.T) {
	type args struct {
		start string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}

	//todo how to mock or setup the fs properly.

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sopsyaml.FindConfigFile(tt.args.start)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindConfigFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindConfigFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadConfigFile(t *testing.T) {
	type args struct {
		confPath string
	}

	tests := []struct {
		name    string
		args    args
		want    *yaml.MapSlice
		wantErr bool
	}{
	// TODO:
	//no file
	//bad yaml
	//good file
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sopsyaml.LoadConfigFile(tt.args.confPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfigFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfigFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteConfigFile(t *testing.T) {
	type args struct {
		confPath string
		yamlMap  *yaml.MapSlice
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO:
	//cant unmarshal
	//cant write
	//writes
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sopsyaml.WriteConfigFile(tt.args.confPath, tt.args.yamlMap); (err != nil) != tt.wantErr {
				t.Errorf("WriteConfigFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
	// TODO:
	//not an array
	//not a string
	//doesnt exist
	//empty
	//stuff
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sopsyaml.ExtractConfigEncryptFiles(tt.args.data)
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
			got, err := sopsyaml.GetConfigEncryptFiles(tt.args.basePath)
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

	tests := []struct {
		name    string
		args    args
		want    *yaml.MapSlice
		wantErr bool
	}{
	// TODO:
	// replaces existing
	// replaces with empty
	// replaces an empty
	// adds new item
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sopsyaml.ReplaceConfigEncryptFiles(tt.args.data, tt.args.encFiles)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReplaceConfigEncryptFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
			if err := sopsyaml.WriteEncryptFilesToDisk(tt.args.confPath, tt.args.data, tt.args.encFiles); (err != nil) != tt.wantErr {
				t.Errorf("WriteEncryptFilesToDisk() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
