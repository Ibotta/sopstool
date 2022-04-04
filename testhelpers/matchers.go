package testhelpers

import (
	"fmt"
	"regexp"

	"github.com/golang/mock/gomock"
)

// RegexMatches creates a Matcher that matches a regex string
// and a value regular expression.
func RegexMatches(valueRegex string) gomock.Matcher {
	return regexMatches{
		valueRegex: regexp.MustCompile(valueRegex),
	}
}

type regexMatches struct {
	valueRegex *regexp.Regexp
}

func (m regexMatches) Matches(x interface{}) bool {
	return m.valueRegex.MatchString(x.(string))
}

func (m regexMatches) String() string {
	return fmt.Sprintf("regexMatches(~/%s/)", m.valueRegex.String())
}
