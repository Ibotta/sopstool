// Copyright © 2017 Ibotta

package sopstool

import "github.com/Ibotta/sopstool/sopsyaml"

// Interface Main sopstool functions
type Interface interface {
	Add(file string, skipEnc bool, skipClean bool) error
	Cat(file string) error
	Clean(file string) error
	Decrypt(file string) error
	Edit(file string) error
	Entrypoint(command string, exec bool, files []string, envfile string) error
	List() error
	Remove(file string, skipClean bool) error
	Rotate(file string) error
}

// SopsTool instance container?
type SopsTool struct {
	cfgPath string
	config  sopsyaml.SopsConfig
}

// NewSopsTool Create a new Sopstool instance
func NewSopsTool(cfgPath string) Interface {
	t := new(SopsTool)
	t.cfgPath = cfgPath
	return t
}

// Add a file
func (st SopsTool) Add(file string, skipEnc bool, skipClean bool) error {

	return nil
}

// Cat prints out a file
func (st SopsTool) Cat(file string) error { return nil }

// Clean a file from the filesystem
func (st SopsTool) Clean(file string) error { return nil }

// Decrypt and save a file
func (st SopsTool) Decrypt(file string) error { return nil }

// Edit a file
func (st SopsTool) Edit(file string) error { return nil }

// Entrypoint calls a command after decrypting
func (st SopsTool) Entrypoint(command string, exec bool, files []string, envfile string) error {
	return nil
}

// List all the files in the config
func (st SopsTool) List() error { return nil }

// Remove a file
func (st SopsTool) Remove(file string, skipClean bool) error { return nil }

// Rotate a file's keys
func (st SopsTool) Rotate(file string) error { return nil }
