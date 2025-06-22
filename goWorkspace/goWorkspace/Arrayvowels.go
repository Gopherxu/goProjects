//To count number of vowels in a string

package main

import (
	"fmt"
	"strings"
)

func countVowels(str string) int {
	count := 0
	vowels := "aeiouAEIOU"

	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}

func main() {
	text := "Hello!! How are you?."
	vowelCount := countVowels(text)
	fmt.Printf("Number of vowels in '%s': %d\n", text, vowelCount)
}