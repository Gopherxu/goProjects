package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func vcount(input string) int {
	c := 0
	vowel := "aeiouAEIOU"
	for _, i := range input {
		if strings.ContainsRune(vowel, i) {
			c++
		}
	}
	return c
}

func ccount(input string) int {
	vowel := "aeiouAEIOU"
	c := 0
	for _, i := range input {
		if !(strings.ContainsRune(vowel, i)) {
			c++
		}
	}
	return c
}

func scount(input string) int {
	s := ' '
	c := 0
	for _, i := range input {
		if i == s {
			c++
		}
	}
	return c
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the string : ")
	input := ""
	if scanner.Scan() {
		input = scanner.Text()
	}

	fmt.Println("This is vowels count : ", vcount(input))
	fmt.Println("This is Consonant count : ", ccount(input)-scount(input))
	fmt.Println("This is Space Count : ", scount(input))
	fmt.Println("This is words count : ", scount(input)+1)
}
