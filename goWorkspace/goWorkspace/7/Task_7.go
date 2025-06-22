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
		log.Fatal("An error occured", err)
	}

	fmt.Println("Enter 2nd number : ")
	fmt.Scanf("%v", &num2)
	n2, err := strconv.Atoi(num2)
	if err != nil {
		log.Fatal("An error occured ", err)
	}

	fmt.Println("Enter operation to be Performed (+,-,*,/) : ")
	fmt.Scanf("%v", &operator)

	var result int
	switch operator {
	case "+":
		result = n1 + n2

	case "-":
		result = n1 - n2

	case "*":
		result = n1 * n2

	case "/":
		if n2 == 0 {
			log.Fatal("num2 Can't be Zero ")
		}
		result = n1 / n2

	default:
		log.Fatal("Enter correct 0perator ")
	}

	fmt.Printf("\n1st num is %v\n2nd num is %v\nThe Operation to be Performed was %v %v %v = %v\n", n1, n2, n1, operator, n2, result)

}
