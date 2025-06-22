package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	var num1, num2 string
	fmt.Printf("Enter 1st Number : ")
	fmt.Scanf("%v", &num1)

	Num1, err := strconv.Atoi(num1) 
	if err != nil {
		log.Fatal("An error Occuted ", err)
	}
	fmt.Printf("Enter 2nd Number : ")
	fmt.Scanf("%v", &num2)

	Num2, err := strconv.Atoi(num2)
	if err != nil {
		log.Fatal("Error in 2nd number ", err)
	}
	fmt.Printf("Num1 = %v and Num2 = %v \n Sum = %v", Num1, Num2, Num1+Num2)

}
