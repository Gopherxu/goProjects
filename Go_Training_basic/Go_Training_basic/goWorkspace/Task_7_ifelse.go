package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var num1, num2, operator string

	fmt.Println("Enter 1st number : ")
	fmt.Scanf("%v", &num1)
	n1, err := strconv.Atoi(num1)
	if err != nil {
		log.Fatal("An error occured ! ")
	}

	fmt.Println("Enter 2Nd number : ")
	fmt.Scanf("%v", &num2)
	n2, err := strconv.Atoi(num2)
	if err != nil {
		log.Fatal("Error occured!!")
	}

	fmt.Println("Enter the operation (+,-,*,/)")
	fmt.Scanf("%v", &operator)

	var result int
	if operator == "+" {
		result = n1 + n2
	} else if operator == "-" {
		result = n1 - n2
	} else if operator == "*" {
		result = n1 * n2

	} else if operator == "/" {
		if n2 == 0 {
			log.Fatal("num2 Can't be Zero ")
		}
		result = n1 / n2
	} else {
		log.Fatal("Entered Operator was wrong !!!")
	}

	fmt.Printf("\nThe result of \n\t%v %v %v = %v\n\n", n1, operator, n2, result)

}
