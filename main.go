// Copyright © 2017 Ibotta

package main

import "github.com/Ibotta/sopstool/cmd"

// Goreleaser default is `-s -w -X main.version={{.Version}} -X main.commit={{.Commit}} -X main.date={{.Date}}`.
var (
	version = "main"
	commit  = "dirty"
	date    = "Now"
)

func main() {
	cmd.BuildVersion = version
	cmd.BuildCommit = commit
	cmd.BuildDate = date
	cmd.Execute()
}
