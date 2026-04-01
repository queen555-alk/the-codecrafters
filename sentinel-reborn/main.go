package main

import (
	"log"
	"os"
	"strings"
	"the-codecrafters/sentinel-reborn/functions"
)

func CallFiles(text string) string {
	input := functions.ToUpper(text)
	input = functions.Cap(input)
	input = functions.Lower(input)
	input = functions.CommandN(input)
	input = functions.Bin(input)
	input = functions.Hex(input)
	input = functions.Vowels(input)
	input = functions.Punctuation(input)
	input = functions.SingleQuotes(input)

	return input
}
func ReadLines(text string) string {
	words := strings.Split(text, "\n")

	for i, line := range words {
		words[i] = CallFiles(line)
	}
	return strings.Join(words, "\n")
}
func main() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := ReadLines(string(data))

	err = os.WriteFile("output.txt", []byte(result), 0644)
	if err != nil {
		log.Fatal(err)
	}

}
