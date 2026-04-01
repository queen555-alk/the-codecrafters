package functions

import (
	"strings"
)

func Vowels(input string) string {
	text := strings.Fields(input)

	val1 := "aeiouhAEIOUH"

	for i, word := range text {
		if string(input[i]) == "A" && strings.ContainsAny(string(word[0]), val1) {
			text[i] = "An"
		}
	}
	return strings.Join(text, " ")
}
