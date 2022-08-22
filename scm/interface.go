package scm

type SCM interface {
	AddFileToIgnored(fn string) error
	RemoveFileFromIgnored(fn string) error
}
