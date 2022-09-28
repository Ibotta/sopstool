package scm

import (
	"os"
	"reflect"
	"testing"
)

func TestRemoveFileFromIgnored(t *testing.T) {
	tests := []struct {
		name             string
		gitIgnoreContent []byte
		fileName         string
		want             []byte
	}{
		{name: "Empty file", gitIgnoreContent: []byte(""), fileName: "secret.yaml", want: []byte("")},
		{name: "Remove from beginning", gitIgnoreContent: []byte("remove.me\nsecret.yml\nvendor"), fileName: "remove.me", want: []byte("secret.yml\nvendor\n")},
		{name: "Remove from the middle", gitIgnoreContent: []byte("secret.yml\nremove.me\nvendor"), fileName: "remove.me", want: []byte("secret.yml\nvendor\n")},
		{name: "Remove from the end ", gitIgnoreContent: []byte("secret.yml\nvendor\nremove.me"), fileName: "remove.me", want: []byte("secret.yml\nvendor\n")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempFilePath := t.TempDir() + "/.gitignore"
			err := os.WriteFile(tempFilePath, tt.gitIgnoreContent, 0644)

			if err != nil {
				t.Errorf("RemoveFileFromIgnored() err=%v", err)
			}

			git := Git{
				IgnoreFilePath: tempFilePath,
			}
			err = git.RemoveFileFromIgnored(tt.fileName)
			if err != nil {
				t.Errorf("RemoveFileFromIgnored() err=%v", err)
			}

			got, err := os.ReadFile(tempFilePath)
			if err != nil {
				t.Errorf("RemoveFileFromIgnored() err=%v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFileFromIgnored() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddFileToIgnored(t *testing.T) {
	tests := []struct {
		name             string
		gitIgnoreContent []byte
		fileName         string
		want             []byte
	}{
		{name: "Empty file", gitIgnoreContent: []byte(""), fileName: "secret.yaml", want: []byte("secret.yaml")},
		{name: "Skip if file exists", gitIgnoreContent: []byte("secret.yml\nvendor"), fileName: "secret.yml", want: []byte("secret.yml\nvendor")},
		{name: "Add file", gitIgnoreContent: []byte("test.txt\nvendor"), fileName: "secret.yml", want: []byte("test.txt\nvendor\nsecret.yml")},
		{name: "Add file with path", gitIgnoreContent: []byte("test.txt\nvendor"), fileName: "../../../secret.yml", want: []byte("test.txt\nvendor\n../../../secret.yml")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempFilePath := t.TempDir() + "/.gitignore"
			err := os.WriteFile(tempFilePath, tt.gitIgnoreContent, 0644)

			if err != nil {
				t.Errorf("AddFileToIgnored() err=%v", err)
			}

			git := Git{
				IgnoreFilePath: tempFilePath,
			}

			err = git.AddFileToIgnored(tt.fileName)
			if err != nil {
				t.Errorf("AddFileToIgnored() err=%v", err)
			}

			got, err := os.ReadFile(tempFilePath)
			if err != nil {
				t.Errorf("AddFileToIgnored() err=%v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFileToIgnored() = %v, want %v", got, tt.want)
			}
		})
	}
}
