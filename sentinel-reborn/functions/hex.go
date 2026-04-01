package functions

import (
	"log"
	"strconv"
	"strings"
)

func Hex(words string) string {
	input := strings.Fields(words)

	for i := 1; i < len(input); i++ {
		if input[i] == "(hex)" && i > 0 {
			val, err := strconv.ParseInt(input[i-1], 16, 64)
			if err != nil {
				log.Fatal(err)
				continue
			}
			input[i-1] = strconv.FormatInt(val, 10)
			input = append(input[:i], input[i+1:]...)
			i--
		}
	}
	return strings.Join(input, " ")
}
