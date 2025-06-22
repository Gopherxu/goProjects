package main

import (
	"fmt"
	"log"
	"time"
)

func holidayList() {
	fmt.Println("Enter dates in holiday list (enter done when completed): ")

	var holidays []time.Time
	for {
		var str string
		fmt.Scanf("%v", &str)
		if str == "done" {
			break
		}

		date, _ := time.Parse("2006-01-02", str)

		holidays = append(holidays, date)
	}

}
func main() {

	var str string
	fmt.Println("Enter Date (YYYY-MM-DD):- ")
	fmt.Scanf("%v", &str)

	date, err := time.Parse("2006-01-02", str)
	if err != nil {
		log.Fatal(err)
	}

	holidayList()

	var days int
	fmt.Println("Enter working days to be skipped  : ")
	fmt.Scan(&days)

	if days > 0 {
		for days > 0 {
			weekday := date.Weekday()

			if weekday != time.Saturday || weekday != time.Sunday {

				date = date.AddDate(0, 0, 1)
				days--
			} else {
				date = date.AddDate(0, 0, 1)
			}
		}

	} else {
		for days < 0 {
			weekday := date.Weekday()

			if weekday != time.Saturday || weekday != time.Sunday {

				date = date.AddDate(0, 0, -1)
				days--
			} else {
				date = date.AddDate(0, 0, -1)
			}
		}

	}

	fmt.Println(date)

}
