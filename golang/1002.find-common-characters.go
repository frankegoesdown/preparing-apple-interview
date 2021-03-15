package main

import (
	"fmt"
)

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
}

func commonChars(A []string) []string {
	commonMap := make(map[rune]int)

	for _, char := range A[0] {
		commonMap[char]++
	}

	for _, words := range A[1:] {

		currentWordsChar := make(map[rune]int)

		for _, char := range words {
			currentWordsChar[char]++
		}

		for char, v := range commonMap {
			if currentWordsChar[char] < v {
				commonMap[char] = currentWordsChar[char]
			}
		}
	}

	var result []string
	for char, v := range commonMap {
		for i := 0; i < v; i++ {
			result = append(result, string(char))
		}
	}

	return result
}
