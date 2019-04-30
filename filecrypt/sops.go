package filecrypt

import (
	"github.com/Ibotta/sopstool/fileutil"
	"github.com/Ibotta/sopstool/oswrap"
)

type sopsCrypt struct {
	execWrap oswrap.ExecWrap
	osWrap   oswrap.OsWrap
}

var sops = sopsCrypt{
	execWrap: oswrap.ExecWrapInstance(),
	osWrap:   oswrap.OsWrapInstance(),
}

// SopsCryptInstance gets an instance of the sops wrapper
func SopsCryptInstance() FileCrypt {
	return sops
}

// EncryptFile encrypts a file rewriting the sops encrypted file
func (sops sopsCrypt) EncryptFile(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)
	return sops.execWrap.RunCommandStdoutToFile(cryptfile, []string{"sops", "-e", fn})
}

// DecryptFile decrypts a file rewriting the plaintext file
func (sops sopsCrypt) DecryptFile(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)
	return sops.execWrap.RunCommandStdoutToFile(fn, []string{"sops", "-d", cryptfile})
}

// DecryptFilePrint decrypts a file printing the result
func (sops sopsCrypt) DecryptFilePrint(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)
	return sops.execWrap.RunCommandDirect([]string{"sops", "-d", cryptfile})
}

// RemoveFile removes a plaintext file from the filesystem
func (sops sopsCrypt) RemoveFile(fn string) error {
	return sops.osWrap.Remove(fn)
}

// RemoveCryptFile removes a sops file from the filesystem
func (sops sopsCrypt) RemoveCryptFile(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)

	return sops.osWrap.Remove(cryptfile)
}

// RotateFile rotates keys on a file
func (sops sopsCrypt) RotateFile(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)

	return sops.execWrap.RunCommandDirect([]string{"sops", "-i", "-r", cryptfile})
}

// EditFile should open the editor for a file
func (sops sopsCrypt) EditFile(fn string) error {
	cryptfile := fileutil.NormalizeToSopsFile(fn)

	return sops.execWrap.RunCommandDirect([]string{"sops", cryptfile})
}
