package functions

import (
	"strings"
)

func Punctuation(word string) string {
	words := strings.Fields(word)
	var result []string

	for _, w := range words {

		for len(w) > 0 && strings.ContainsAny(string(w[0]), "!.,;?") {
			if len(result) > 0 {

				result[len(result)-1] += string(w[0])
			}
			w = w[1:]
		}

		result = append(result, w)

	}

	return strings.Join(result, " ")
}
