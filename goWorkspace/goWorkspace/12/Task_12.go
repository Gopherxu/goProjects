package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var num1, num2 string
	fmt.Println("Enter Two numbers : ")
	fmt.Scanf("%v %v", &num1, &num2)

	n1, err := strconv.Atoi(num1)
	if err != nil {
		log.Fatal("1st Value Entered Was not a Number !!", err)
	}
	n2, err := strconv.Atoi(num2)
	if err != nil {
		log.Fatal("2nd Value Entered is Not an Number !!", err)
	}
	for i := 1; i <= n2; i++ {
		fmt.Printf("%v * %v = %v\n", n1, i, n1*i)
	}

}
