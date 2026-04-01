package functions

import (
	"regexp"
)

func SingleQuotes(s string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)

	return re.ReplaceAllString(s, "'$1'")
}
