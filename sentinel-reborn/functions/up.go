package main

import (
	"strconv"
	"strings"
)

func ToUpperCase(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}

		if strings.HasPrefix(words[i], "(up,") && i > 0 {
			numStr := strings.TrimSuffix(words[i+1], ")")
			n, _ := strconv.Atoi(numStr)
			for j := 0; j < n && i-1-j >= 0; j++ {
				words[i-1-j] = strings.ToUpper(words[i-1-j])
			}

			words = append(words[:i], words[i+2:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}
