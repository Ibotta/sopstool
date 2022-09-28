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

func TestRemoveFileNoExistingFile(t *testing.T) {
	t.Run("missing gitignore file", func(t *testing.T) {
		tempFilePath := t.TempDir() + "/.gitignore"
		if _, err := os.Stat(tempFilePath); err == nil {
			t.Errorf("RemoveFileFromIgnored() = gitignore accidentally exists")
		}

		git := Git{
			IgnoreFilePath: tempFilePath,
		}
		err := git.RemoveFileFromIgnored("secret.yaml")
		if err != nil {
			t.Errorf("RemoveFileFromIgnored() err=%v", err)
		}

		if _, err := os.Stat(tempFilePath); err == nil {
			t.Errorf("RemoveFileFromIgnored() = gitignore exists when it should not")
		}
	})
}

func TestAddFileToIgnored(t *testing.T) {
	tests := []struct {
		name             string
		gitIgnoreContent []byte
		fileName         string
		want             []byte
	}{
		{name: "Empty file", gitIgnoreContent: []byte(""), fileName: "secret.yaml", want: []byte("secret.yaml\n")},
		{name: "Skip if file exists", gitIgnoreContent: []byte("secret.yml\nvendor"), fileName: "secret.yml", want: []byte("secret.yml\nvendor")},
		{name: "Add file to end", gitIgnoreContent: []byte("test.txt\nvendor"), fileName: "secret.yml", want: []byte("test.txt\nvendor\nsecret.yml\n")},
		{name: "Add file to end, trailing newline", gitIgnoreContent: []byte("test.txt\nvendor\n"), fileName: "secret.yml", want: []byte("test.txt\nvendor\nsecret.yml\n")},
		{name: "Add file with path", gitIgnoreContent: []byte("test.txt\nvendor"), fileName: "../../../secret.yml", want: []byte("test.txt\nvendor\n../../../secret.yml\n")},
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

func TestAddFileNoExistingFile(t *testing.T) {
	t.Run("no existing gitignore", func(t *testing.T) {
		tempFilePath := t.TempDir() + "/.gitignore"
		if _, err := os.Stat(tempFilePath); err == nil {
			t.Errorf("RemoveFileFromIgnored() = gitignore accidentally exists")
		}

		git := Git{
			IgnoreFilePath: tempFilePath,
		}
		err := git.AddFileToIgnored("secret.yaml")
		if err != nil {
			t.Errorf("AddFileToIgnored() err=%v", err)
		}

		got, err := os.ReadFile(tempFilePath)
		if err != nil {
			t.Errorf("AddFileToIgnored() err=%v", err)
		}

		want := []byte("secret.yaml\n")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AddFileToIgnored() = %v, want %v", got, want)
		}
	})
}
