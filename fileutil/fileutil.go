package fileutil

import (
	"fmt"
	"path"
	"regexp"
	"strings"
)

const (
	sopsCryptedFileSegment = ".sops"
)

var sopsFileRegex = regexp.MustCompile(`(.*)\.sops(\.\w+)?$`)

// NormalizeToPlaintextFile normalizes a filename string to not include the 'sops' tag to represent the plaintext version
func NormalizeToPlaintextFile(fn string) string {
	//todo this should normalize it to be relative to the configfile path
	if sopsFileRegex.MatchString(fn) {
		fn = sopsFileRegex.ReplaceAllString(fn, "$1$2")
	}

	return fn
}

// NormalizeToSopsFile normalizes a filename string to include the 'sops' tag to represent the crypted version
func NormalizeToSopsFile(fn string) string {
	//todo this should normalize it to be relative to the configfile path
	if sopsFileRegex.MatchString(fn) {
		return fn
	}

	ext := path.Ext(fn)

	if len(ext) > 0 {
		fn = strings.Replace(fn, ext, sopsCryptedFileSegment+ext, 1)
	} else {
		fn = fn + sopsCryptedFileSegment
	}

	return fn
}

// ListIndexOf gets index of element in list, or -1
func ListIndexOf(files []string, fn string) int {
	found := -1
	fn = NormalizeToPlaintextFile(fn) //TODO is this always redundant?
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
