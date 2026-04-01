package functions

import (
	"log"
	"strconv"
	"strings"
)

func BinToDecimal(str string) string {
	slice := strings.Fields(str)

	for i := 1; i < len(slice); i++ {
		if slice[i] == "(bin)" && i > 0 {
			con, err := strconv.ParseInt(slice[i-1], 2, 64)
			if err != nil {
				log.Fatal(err)
				continue
			}
			slice[i-1] = strconv.FormatInt(con, 10)
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}

	}
	return strings.Join(slice, " ")
}
