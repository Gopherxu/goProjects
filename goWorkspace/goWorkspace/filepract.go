package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	filename := "Doc.txt"

	// to create file
	file, _ := os.Create(filename)

	defer file.Close()
	fmt.Println("File Created : ", file)

	//to add content to file
	str := "\nHey, i have created my first file in go\n "
	file.WriteString(str)

	fmt.Println("Content added to file", filename)

	//Readig content of file

	data, _ := ioutil.ReadFile(filename)
	//fmt.Println("This is file content : ", string(data))

	//Additional content

	additionalContent := "This is additional content.\n"
	file.WriteString(additionalContent)

	data, _ = ioutil.ReadFile(filename)

	fmt.Println("This is file content : ", string(data))

	//to get word count

	words := strings.Fields(string(data))
	wordscount := len(words)
	fmt.Println("\n", wordscount)

}
