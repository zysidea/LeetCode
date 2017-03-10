package main

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{"Hello", "Alaska", "Dad", "Peace"}
	output := findWords(input)
	for _, v := range output {
		fmt.Print(v + " ")
	}
}

var (
	firstKey = []string{"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P", "q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}

	secondKey = []string{"A", "S", "D", "F", "G", "H", "J", "K", "L", "a", "s", "d", "f", "g", "h", "j", "k", "l"}

	thirdKey = []string{"Z", "X", "C", "V", "B", "N", "M", "z", "x", "c", "v", "b", "n", "m"}
)

func findWords(words []string) []string {
	var result []string
	for _, v1 := range words {
		match1 := match(v1, firstKey)
		match2 := match(v1, secondKey)
		match3 := match(v1, thirdKey)

		if match1 && !match2 && !match3 {
			result = append(result, v1)
		} else if !match1 && match2 && !match3 {
			result = append(result, v1)
		} else if !match1 && !match2 && match3 {
			result = append(result, v1)
		} else {
			continue
		}
	}
	return result
}

func match(str string, matchs []string) bool {
	for _, v := range matchs {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}
