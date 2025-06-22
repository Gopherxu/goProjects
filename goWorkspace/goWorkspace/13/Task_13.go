package main

import "fmt"

func main() {
	var char string
	var num1, num2 int
	fmt.Println("Enter charactor, Row size and Collumn Size : ")
	fmt.Scanf("%v %v %v", &char, &num1, &num2)
	for i := 0; i < num1; i++ {
		for j := 0; j < num2; j++ {
			fmt.Printf("%v ", char)
		}
		println()
	}
}
