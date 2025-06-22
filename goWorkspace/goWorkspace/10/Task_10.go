package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var arr []int
	var str string
	var sum int = 0
	fmt.Println(`Enter Numbers or Else Enter "Proceed" to get sorted numbers and their sum : `)
	for {
		fmt.Scanf("%v", &str)
		if strings.ToLower(str) == "proceed" {
			break
		} else {
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal("Enter value was not an integer - ", err)
			}
			sum = sum + num

			arr = append(arr, num)
		}
	}
	sort.Ints(arr)
	fmt.Println("Sum of given elements is : ", sum)
	fmt.Println(arr)
}
