package functions

import (
	"strconv"
	"strings"
)

func CommandN(input string) string {
	text := strings.Fields(input)

	for i := 0; i < len(text); i++ {
		//for (cap, N)
		if strings.HasPrefix(text[i], "(cap,") {
			var Num string
			if strings.HasSuffix(text[i], ")") {
				Num = strings.TrimSuffix(strings.TrimPrefix(text[i], "(cap,"), ")")
				text = append(text[:i], text[i+1:]...)
				i--
			}
			if i+1 < len(text) {
				Num = strings.TrimSuffix(text[i+1], ")")
				text = append(text[:i], text[i+2:]...)
				i--
			}
			//convert number to alpha
			count, _ := strconv.Atoi(Num)

			//effect point
			start := i - count + 1
			if start < 0 {
				start = 0
			}

			for j := start; j <= i; j++ {
				text[j] = strings.ToUpper(string(text[j][0])) + strings.ToLower(text[j][1:])
			}
		}
		//for (low, N)
		if strings.HasPrefix(text[i], "(low,") {
			var Num string
			if strings.HasSuffix(text[i], ")") {
				Num = strings.TrimSuffix(strings.TrimPrefix(text[i], "(low,"), ")")
				text = append(text[:i], text[i+1:]...)
				i--
			}
			if i+1 < len(text) {
				Num = strings.TrimSuffix(text[i+1], ")")
				text = append(text[:i], text[i+2:]...)
				i--
			}
			//convert number to alpha
			count, _ := strconv.Atoi(Num)

			start := i - count + 1
			if start < 0 {
				start = 0
			}

			for j := start; j <= i; j++ {
				text[j] = strings.ToLower(text[j])
			}
		}
		//for (low, N)
		if strings.HasPrefix(text[i], "(up,") {
			var Num string
			if strings.HasSuffix(text[i], ")") {
				Num = strings.TrimSuffix(strings.TrimPrefix(text[i], "(up"), ")")
				text = append(text[:i], text[i+1:]...)
				i--
			}
			if i+1 < len(text) {
				Num = strings.TrimSuffix(text[i+1], ")")
				text = append(text[:i], text[i+2:]...)
				i--
			}
			count, _ := strconv.Atoi(Num)

			start := i - count + 1
			if start < 0 {
				start = 0
			}

			for j := start; j <= i; j++ {
				text[j] = strings.ToUpper(text[j])
			}
		}
	}
	return strings.Join(text, " ")
}
