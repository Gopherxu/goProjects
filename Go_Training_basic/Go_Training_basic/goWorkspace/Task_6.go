package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Println("Enter 1st number : ")
	fmt.Scanf("%v", &num1)

	// n1, err := strconv.ParseFloat(num1, 32)
	// if err != nil {
	// 	log.Fatal("An error in 1st num ", err)
	// }

	fmt.Println("Enter 2nd num : ")
	fmt.Scanf("%v", &num2)

	// n2, err := strconv.ParseFloat(num2, 32)
	// if err != nil {
	// 	log.Fatal("An error in 2nd numm", err)
	// }

	fmt.Printf("num1 = %v and num2 = %v \nsum = %v", num1, num2, num1+num2)
}
