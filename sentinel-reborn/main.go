package main

import (
	"log"
	"os"
	"strings"
	"the-codecrafters/sentinel-reborn/functions"
)

func CallFiles(text string) string {
	input := functions.Vowels(text)
	input = functions.Bin(input)
	input = functions.Cap(input)
	input = functions.CommandN(input)
	input = functions.Hex(input)
	input = functions.Punctuation(input)
	input = functions.SingleQuotes(input)
	input = functions.ToUpper(input)

	return text
}
func ReadLines(text string) string {
	words := strings.Split(text, "\n")

	for i, line := range words {
		words[i] = CallFiles(line)
	}
	return strings.Join(words, "\n")
}
func main() {
	inputfile := os.Args[1]
	outputfile := os.Args[2]

	data, err := os.ReadFile(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	result := CallFiles(string(data))

	err = os.WriteFile(outputfile, []byte(result), 0644)
	if err != nil {
		log.Fatal(err)
	}

}
