package main

import (
	"fmt"
	"math"
	"strconv"
)

var validNums int = 0

// Count function :- It counts number of Valid inputs
func count() int {

	return validNums
}

func mean(arr []int) float64 {
	var sum int = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(len(arr))

}

func min(arr []int) int {
	var min int = math.MaxInt64
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func max(arr []int) int {
	max := math.MinInt64
	for i := 0; i < len(arr); i++ {
		if int(arr[i]) > int(max) {
			max = int(arr[i])
		}
	}
	return max
}

func codd(arr []int) int {
	c := 0
	for _, i := range arr {
		if i%2 == 1 {
			c++
		}
	}
	return c
}

func ceven(arr []int) int {
	c := 0
	for _, i := range arr {
		if i%2 == 0 {
			c++
		}
	}
	return c
}

func main() {

	var str string
	var arr []int

	fmt.Println("Enter numbers or Enter Operations you want to Perform (count, mean, min, max,odd_nums,even_nums)")
	for {
		fmt.Scanf("%v", &str)
		if (str == "count") || (str == "min") || (str == "mean") || (str == "max") || (str == "even_nums") || (str == "odd_nums") {
			break
		}
		n, err := strconv.Atoi(str)
		if err != nil {

		} else {
			arr = append(arr, n)
			validNums++
		}

	}
	switch str {
	case "count":
		fmt.Println("The count of numbers Entered is : ", count())

	case "mean":
		fmt.Println("The Mean of Given NUmbers is : ", mean(arr))

	case "min":
		fmt.Println("Minimum of All numbers is : ", min(arr))

	case "max":
		fmt.Println("Maximum of All Numbers is : ", max(arr))

	case "odd_nums":
		fmt.Println("Count of Odd Numbers Entered is : ", codd(arr))

	case "even_nums":
		fmt.Println("Count of Even Numbres Entered is : ", ceven(arr))
	}
}
