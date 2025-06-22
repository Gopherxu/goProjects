// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func randomNum() int {
// 	randomNumber := rand.Intn(100) + 1
// 	fmt.Println(randomNumber)
// 	return randomNumber
// }
// func main() {
// 	var arr []int
// 	for len(arr) <= 99 {
// 		num := randomNum()
// 		var flag int = 0
// 		for _, value := range arr {
// 			if value == num {
// 				flag = 1
// 			}
// 		}
// 		if flag == 0 {
// 			arr = append(arr, num)
// 		}
// 	}
// 	fmt.Println(arr)
// 	fmt.Println(len(arr))

// }

package main

import (
	"fmt"
	"time"
)

func main() {
	// Get user input for start and end dates
	var startDateStr, endDateStr string
	fmt.Println("Enter start date (YYYY-MM-DD): ")
	fmt.Scanln(&startDateStr)
	fmt.Println("Enter end date (YYYY-MM-DD): ")
	fmt.Scanln(&endDateStr)

	// Parse the start and end dates
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		fmt.Println("Error parsing start date:", err)
		return
	}
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		fmt.Println("Error parsing end date:", err)
		return
	}

	// Count weekend days

	var holiday int
	for d := startDate; d.Before(endDate); d = d.AddDate(0,0,1) {

		weekdays := startDate.Weekday()
		if weekdays == time.Saturday || weekdays == time.Sunday {
			holiday++
		}
		startDate = startDate.AddDate(0, 0, 1)
	}
	fmt.Println("Number of weekend days between", startDateStr, "and", endDateStr, "is", holiday)

}
