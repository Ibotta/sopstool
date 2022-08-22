package scm

import (
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Git struct {
}

// AddFileToIgnored adds filename to .gitignore
func (g Git) AddFileToIgnored(fn string) error {
	exists, err := lineInFileExists(fn, ".gitignore")

	if err != nil {
		return err
	}

	if !exists {
		return appendLineToFileIfNotExists(fn, ".gitignore")
	} else {
		fmt.Println("File already exists in .gitignore file. Skipping.")
		return nil
	}
}

// RemoveFileFromIgnored removes filename from .gitignore
func (g Git) RemoveFileFromIgnored(fn string) error {
	exists, err := lineInFileExists(fn, ".gitignore")

	if err != nil {
		return err
	}
	if !exists {
		fmt.Println("file not exists in .gitignore file. Skipping.")
		return nil
	} else {
		return removeLineFromFile(fn, ".gitignore")
	}
}

// lineInFileExists verifies if line exists in provided file
func lineInFileExists(line string, filename string) (bool, error) {
	b, err := os.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		} else if errors.Is(err, os.ErrPermission) {
			return false, fmt.Errorf("insufficient permissions to read %s file", filename)
		} else {
			return false, err
		}
	}
	// Regex pattern captures line from the content.
	match, _ := regexp.Match("(?m)^"+regexp.QuoteMeta(line)+"$", b)

	return match, nil
}

// appendLineToFileIfNotExists appends line to file if not exists. It also creates file if not exists
func appendLineToFileIfNotExists(line string, filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	// Verify if last char is a new line
	if string(b[len(b)-1]) != "\n" {
		line = "\n" + line
	}

	if _, err := file.WriteString(line); err != nil {
		return err
	}
	return nil
}

// removeLineFromFile removes provided line from filename
func removeLineFromFile(line string, filename string) error {
	b, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	newFileContent := strings.Replace(string(b), "\n"+line, "", -1)

	err = os.WriteFile(filename, []byte(newFileContent), 0)
	if err != nil {
		return err
	}
	return nil
}
