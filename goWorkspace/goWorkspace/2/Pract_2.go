package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func ecount(nums []int) int {
	c := 0
	for _, i := range nums {
		if i%2 == 0 {
			c++
		}
	}
	return c
}

func ocount(nums []int) int {
	c := 0
	for _, i := range nums {
		if i%2 == 1 {
			c++
		}
	}
	return c
}

func main() {
	var arr []string
	var nums []int
	var str string
	i := 0
	fmt.Println("Enter the elements in Array or else Enter Proceed : ")
	for {
		fmt.Scanf("%v", &str)
		arr = append(arr, str)
		if strings.ToLower(arr[i]) == "proceed" {
			break
		} else {
			num, err := strconv.Atoi(arr[i])
			if err != nil {
				log.Fatal(`Enterd value is not a number neihter "Proceed" - `, err)
			}
			nums = append(nums, num)
		}
		i++

	}
	fmt.Println("The Count of Even Numbers in Array is : ", ecount(nums))
	fmt.Println("The Count of Odd  Numbers is Array is : ", ocount(nums))
}
