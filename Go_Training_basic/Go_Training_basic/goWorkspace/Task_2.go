package main

import (
	"fmt"
	"log"
)

func main() {
	var name string
	fmt.Printf("Enter Your Name :- ")
	fmt.Scan(&name)

	if name == " " {
		log.Fatal("Enter Valid name !!!!!")
	}

	fmt.Println("Hello, ", name)
}
