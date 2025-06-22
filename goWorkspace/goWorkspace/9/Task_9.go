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

func main() {

	var str string
	var arr []int

	fmt.Println("Enter numbers or Enter Operations you want to Perform (count, mean, min, max)")
	for {
		fmt.Scanf("%v", &str)
		if (str == "count") || (str == "min") || (str == "mean") || (str == "max") {
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
	}
}
