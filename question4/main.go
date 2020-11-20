package main

import (
	"fmt"
	"sort"
)

type sortRunes []rune

var strArr = []string{
	"kita",
	"atik",
	"tika",
	"aku",
	"kia",
	"makan",
	"kua",
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

// SortString : convert string to rune type then sorting it
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// FindRepeatitionWord : convert same word into map
func FindRepeatitionWord(str []string) map[string]int {
	repeat := make(map[string]int)

	for _, word := range str {
		_, matched := repeat[word]
		if matched {
			repeat[word]++
		} else {
			repeat[word] = 1
		}
	}
	return repeat
}

func main() {
	var result []string
	var finalResult [][]string

	// convert string to rune then sorting it
	for i := 0; i < len(strArr); i++ {
		tmp := SortString(strArr[i])
		result = append(result, tmp)
	}

	// grouping same anagram
	strToMap := FindRepeatitionWord(result)

	// mapping to slice
	for key := range strToMap {
		var temp []string
		for i := 0; i < len(strArr); i++ {
			if key == result[i] {
				temp = append(temp, strArr[i])
			}
		}
		finalResult = append(finalResult, temp)

	}

	fmt.Println(finalResult)

}
