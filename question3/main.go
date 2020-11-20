package main

import (
	"fmt"
	"strings"
)

// CONSTANT ERROR
const (
	ErrorEmptyString     = "Empty String!"
	ErrorBracketNotFound = "Bracket Not found!"
)

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1])
			}
		} else {
			return ErrorBracketNotFound
		}
	}
	return ErrorEmptyString
}

func main() {
	var str string
	// var str = "(Hello)" // please enable this line to used fix string
	fmt.Println("Input String : ")
	fmt.Scanf("%s", &str)

	result := findFirstStringInBracket(str)

	fmt.Println("First string in bracket : " + result)
}
