package functions

import (
	"strings"
)

func cap(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {

		if words[i] == "(cap)" && i > 0 {
			w := words[i-1]
			if len(w) > 0 {

				words[i-1] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])

			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")

}
