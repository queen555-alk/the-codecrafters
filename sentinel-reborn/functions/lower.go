package functions

import (
	"strings"
)

// queen is to work here
func Lower(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}
