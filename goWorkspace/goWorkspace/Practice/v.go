package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	var num1, num2 string

	fmt.Println("Enter 1st number : ")
	fmt.Scan(&num1)

	firstnumber, err := strconv.Atoi(num1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter 1st number : ")
	fmt.Scan(&num2)

	secondumber, err := strconv.Atoi(num2)
	if err != nil {
		log.Fatal(err)
	}

	var pattern string

	fmt.Scan(&pattern)

	for i := 0; i < firstnumber; i++ {
		for j := 0; j < secondumber; j++ {
			fmt.Printf("%v ", pattern)
		}
		fmt.Println()
	}
}
