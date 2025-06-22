package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to Open File
func openFile(fileName string) *os.File {

	file, err := os.Open(fileName)

	//Handling Error if any occures
	if err != nil {
		fmt.Println("Error Occured While opeaning File ", err)
	}

	return file
}

// This function Calculates Total word count and count of each indivual word
func wordCount(file *os.File) (int, map[string]int) {

	//Created Scanner to Read line one by one
	scanner := bufio.NewScanner(file)

	//Map to Store word frequency
	WordCount := make(map[string]int)

	//Varible to get count of all Words
	var totalWords int

	//scan each line    "it return true if there is new line else returns false"
	for scanner.Scan() {

		line := scanner.Text()

		//split line into Words
		words := strings.Fields(line)

		//updating word Count and word Frequency
		for _, word := range words {
			word = strings.ToLower(word)

			//To increase frequency of Specific Word
			WordCount[word]++

			totalWords++
		}

	}
	return totalWords, WordCount
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please Enter file Name as an Argument  !! ")
	}

	fileName := os.Args[1]

	//open File to Read its content
	file := openFile(fileName)

	//store word count into totalwords and frequency into wordcount map
	totalWords, WordCount := wordCount(file)

	//Print total word count
	fmt.Println("Count of Total Words in file is : ", totalWords)

	//TO Print frequency of words
	for word, count := range WordCount {
		fmt.Printf("\n %v: %v \n", word, count)
	}

}
