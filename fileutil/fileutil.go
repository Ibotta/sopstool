package fileutil

import (
	"fmt"
	"path"
	"strings"
)

const (
	sopsCryptedFileSegment = ".sops"
)

// NormalizeToPlaintextFile normalizes a filename string to not include the 'sops' tag to represent the plaintext version
func NormalizeToPlaintextFile(fn string) string {
	//todo this should normalize it to be relative to the configfile path
	if strings.Contains(fn, sopsCryptedFileSegment) {
		fn = strings.Replace(fn, sopsCryptedFileSegment, "", 1)
	}

	return fn
}

// NormalizeToSopsFile normalizes a filename string to include the 'sops' tag to represent the crypted version
func NormalizeToSopsFile(fn string) string {
	//todo this should normalize it to be relative to the configfile path
	if strings.Contains(fn, sopsCryptedFileSegment) {
		return fn
	}

	ext := path.Ext(fn)
	fn = strings.Replace(fn, ext, sopsCryptedFileSegment+ext, 1)

	return fn
}

// ListIndexOf gets index of element in list, or -1
func ListIndexOf(files []string, fn string) int {
	found := -1
	fn = NormalizeToPlaintextFile(fn)
	for i, f := range files {
		if f == fn {
			found = i
			break
		}
	}
	return found
}

// SomeOrAllFiles gives all or matching files
func SomeOrAllFiles(args []string, encFiles []string) ([]string, error) {
	filesToReturn := []string{}
	if len(args) > 0 {
		for _, fileArg := range args {
			//find mentioned file and decrypt
			fn := NormalizeToPlaintextFile(fileArg)
			if ListIndexOf(encFiles, fn) < 0 {
				return nil, fmt.Errorf("File not found: %s", fn)
			}
			filesToReturn = append(filesToReturn, fn)
		}
	} else {
		filesToReturn = encFiles
	}

	return filesToReturn, nil
}
