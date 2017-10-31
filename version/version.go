package version

import (
	"fmt"
	"strconv"
	"time"
)

var (
	Number        = "0.1.0"
	CommitHash    = "<undefined>"
	CommitStamp   = "<undefined>"
	BuildStamp    = "<undefined>"
	CiBuildNumber = "<undefined>"

	Long string
)

func init() {
	stamp, _ := strconv.Atoi(CommitStamp)
	commitDate := time.Unix(int64(stamp), 0).UTC().Format("2006/01/02-15:04")
	stamp, _ = strconv.Atoi(BuildStamp)
	buildDate := time.Unix(int64(stamp), 0).UTC().Format("Mon Jan 02 15:04:05 MST 2006")
	Long = fmt.Sprintf(`%s-%s (%s) built #%s at %s`,
		Number, CommitHash, commitDate,
		CiBuildNumber, buildDate,
	)
}
