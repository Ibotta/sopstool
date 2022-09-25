package scm

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

type Git struct {
	IgnoreFilePath string
}

// AddFileToIgnored adds filename to .gitignore
func (g Git) AddFileToIgnored(fn string) error {
	exists, err := lineInFileExists(fn, g.IgnoreFilePath)

	if err != nil {
		return err
	}

	if !exists {
		return appendLineToFileIfNotExists(fn, g.IgnoreFilePath)
	} else {
		fmt.Println("File already exists in .gitignore file. Skipping.")
		return nil
	}
}

// RemoveFileFromIgnored removes filename from .gitignore
func (g Git) RemoveFileFromIgnored(fn string) error {
	exists, err := lineInFileExists(fn, g.IgnoreFilePath)

	if err != nil {
		return err
	}
	if !exists {
		fmt.Println("file not exists in .gitignore file. Skipping.")
		return nil
	} else {
		return removeLineFromFile(fn, g.IgnoreFilePath)
	}
}

// lineInFileExists verifies if line exists in provided file
func lineInFileExists(line string, filename string) (bool, error) {
	file, err := os.Open(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		} else if errors.Is(err, os.ErrPermission) {
			return false, fmt.Errorf("insufficient permissions to read %s file", filename)
		} else {
			return false, err
		}
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Regex pattern captures specific line in file
	r, err := regexp.Compile("^" + regexp.QuoteMeta(line) + "$")

	if err != nil {
		return false, err
	}

	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}

// appendLineToFileIfNotExists appends line to file if not exists. It also creates file if not exists
func appendLineToFileIfNotExists(line string, filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	// Get file info
	fi, err := file.Stat()

	if err != nil {
		return err
	}

	// Verify last char only if file is not empty
	if fi.Size() > 0 {
		// Declare a buffer for storing last file character
		b := make([]byte, 1)

		// Set offset for the next Read to the last char
		_, err = file.Seek(-1, 2)

		if err != nil {
			return err
		}

		// Read the last char
		_, err = file.Read(b)
		if err != nil {
			return err
		}

		// Verify if last char is not a new line
		if len(b) > 0 && string(b) != "\n" {
			line = "\n" + line
		}
	}

	if _, err := file.WriteString(line); err != nil {
		return err
	}
	return nil
}

// removeLineFromFile removes provided line from filename
func removeLineFromFile(line string, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create temp file for storing new content of target file
	tempFile, err := os.CreateTemp("/tmp", "sopstool")
	if err != nil {
		return err
	}

	defer tempFile.Close()
	//Write file conent omitting specific line to tempFile
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != line {
			_, err := tempFile.WriteString(scanner.Text() + "\n")
			if err != nil {
				return err
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// Override target file with tempFile
	err = os.Rename(tempFile.Name(), filename)
	if err != nil {
		return err
	}
	return nil
}
