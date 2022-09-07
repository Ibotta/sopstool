package filecrypt

type FileCrypt interface {
	EncryptFile(fn string) error
	DecryptFile(fn string) error
	DecryptFilePrint(fn string) error
	RemoveFile(fn string) error
	RemoveCryptFile(fn string) error
	RotateFile(fn string) error
	EditFile(fn string) error
	UpdateKeysFile(fn string, extraArgs []string) error
}
