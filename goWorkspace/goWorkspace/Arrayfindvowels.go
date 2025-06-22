//To find the vowels in a string and to print them

package main
import "fmt"

func findVowels(str string) string {
	vowels := ""
	for _, char := range str {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowels += string(char)
		}
	}
	return vowels
}

func main() {
	text := "Hello!! How are you?."
	vowelString := findVowels(text)
	fmt.Printf("Vowels in '%s': %s\n", text, vowelString)
}
