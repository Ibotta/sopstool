package sopsyaml

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/mozilla-services/yaml" //this branch has the unmarshaler that keeps comments
)

// This is all the stuff needed to import the sops.yaml config file

const (
	maxDepth          = 100
	configFileName    = ".sops.yaml"
	stopFileName      = ".git"
	encryptedFilesKey = "encrypted_files"
)

// SopsConfig holds info about an instance of the config file
type SopsConfig struct {
	Path           string
	Tree           *yaml.MapSlice
	EncryptedFiles []string
}

// wrap OS filesystem commands for mocking
type fileSystem interface {
	Stat(name string) (os.FileInfo, error)
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, data []byte, perm os.FileMode) error
}
type osFS struct {
	stat      func(string) (os.FileInfo, error)
	readfile  func(string) ([]byte, error)
	writefile func(string, []byte, os.FileMode) error
}

func (fs osFS) Stat(name string) (os.FileInfo, error) {
	return fs.stat(name)
}
func (fs osFS) ReadFile(filename string) ([]byte, error) {
	return fs.readfile(filename)
}
func (fs osFS) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return fs.writefile(filename, data, perm)
}

var fs fileSystem = osFS{
	stat:      os.Stat,
	readfile:  ioutil.ReadFile,
	writefile: ioutil.WriteFile,
}

// FindConfigFile looks for a sops config file in the current working directory and on parent directories, up to the limit defined by the maxDepth constant.
func FindConfigFile(start string) (string, error) {
	filepath := start
	for i := 0; i < maxDepth; i++ {
		_, err := fs.Stat(path.Join(filepath, configFileName))
		if err != nil {
			_, giterr := fs.Stat(path.Join(filepath, ".git"))
			if giterr == nil {
				//found top of git, stop here
				break
			}
			//next iteration will be one higher
			filepath = path.Join(filepath, "..")
		} else {
			return path.Join(filepath, configFileName), nil
		}
	}
	//TODO gracefully create a file at `git root`
	return "", fmt.Errorf("Config file not found")
}

// LoadConfigFile loads a yaml file path into a yaml map
func LoadConfigFile(confPath string) (*yaml.MapSlice, error) {
	confBytes, err := fs.ReadFile(confPath)
	if err != nil {
		return nil, fmt.Errorf("could not read config file: %s", err)
	}

	var data yaml.MapSlice
	if err := (yaml.CommentUnmarshaler{}).Unmarshal(confBytes, &data); err != nil {
		return nil, fmt.Errorf("Error unmarshaling input YAML: %s", err)
	}

	return &data, nil
}

// WriteConfigFile writes out a yaml file
func WriteConfigFile(confPath string, yamlMap *yaml.MapSlice) error {
	out, err := (&yaml.YAMLMarshaler{Indent: 2}).Marshal(yamlMap)
	if err != nil {
		return fmt.Errorf("Error marshaling to yaml: %s", err)
	}
	fs.WriteFile(confPath, out, 0644)

	return nil
}

// ExtractConfigEncryptFiles pulls the files we want to manipulate out of the map
func ExtractConfigEncryptFiles(data *yaml.MapSlice) ([]string, error) {
	encFiles := []string{}
	for _, item := range *data {
		if item.Key == encryptedFilesKey {
			//assert that this is a slice
			listSlice, ok := item.Value.([]interface{})
			if !ok {
				return nil, fmt.Errorf("encrypted_files is not an array")
			}
			for _, v := range listSlice {
				value, ok := v.(string)
				if !ok {
					return nil, fmt.Errorf("encrypted_files must be an array of strings")
				}
				encFiles = append(encFiles, value)
			}
			break
		}
	}
	return encFiles, nil
}

// GetConfigEncryptFiles is a shortcut for getting the file list when no other list data is required
func GetConfigEncryptFiles(basePath string) ([]string, error) {
	cfgFile, err := FindConfigFile(basePath)
	if err != nil {
		return nil, err
	}
	data, err := LoadConfigFile(cfgFile)
	if err != nil {
		return nil, fmt.Errorf("Error loading config: %s", err)
	}
	encFiles, err := ExtractConfigEncryptFiles(data)
	if err != nil {
		return nil, fmt.Errorf("Error reading config: %s", err)
	}

	return encFiles, nil
}

// ReplaceConfigEncryptFiles pulls the files we want to manipulate out of the map
func ReplaceConfigEncryptFiles(data *yaml.MapSlice, encFiles []string) (*yaml.MapSlice, error) {
	//remake the root data
	out := make(yaml.MapSlice, 0)
	found := false
	for _, item := range *data {
		if item.Key == encryptedFilesKey {
			found = true
			item.Value = encFiles
		}
		out = append(out, item)
	}
	if !found {
		//didnt find an existing encrypted_files element, add it
		out = append(out, yaml.MapItem{Key: encryptedFilesKey, Value: encFiles})
	}
	return &out, nil
}

// WriteEncryptFilesToDisk writes the new files to disk based on existing data
func WriteEncryptFilesToDisk(confPath string, data *yaml.MapSlice, encFiles []string) error {
	outdata, err := ReplaceConfigEncryptFiles(data, encFiles)
	if err != nil {
		return fmt.Errorf("Error replacing config: %s", err)
	}
	err = WriteConfigFile(confPath, outdata)
	if err != nil {
		return fmt.Errorf("Error writing config: %s", err)
	}
	return nil
}
