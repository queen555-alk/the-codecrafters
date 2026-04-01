package functions

import (
	"regexp"
)

func singleQuotes(s string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)

	return re.replaceAllString(s, "'$1'")
}
