package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var str string
	var sum int = 0
	fmt.Println("Enter Numbers or Enter Proceed to get sum : ")

	for {
		fmt.Scanf("%v", &str)
		if str == "proceed" {
			break
		}
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal("Entered Value was't Number !!  ")
		}

		sum += num
	}
	fmt.Printf("\nSum of Numbers Entered is : %v \n\n", sum)
}
