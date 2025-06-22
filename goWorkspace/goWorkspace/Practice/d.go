package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var str string
	fmt.Scanf("%v", &str)

	date, err := time.Parse("2006-01-02", str)
	if err != nil {
		log.Fatal(err)
	}
	var num int
	var count int
	fmt.Scanf("%v", &num)
	if num > 0 {
		for num > 0 {

			weekdate := date.Weekday()
			if weekdate == time.Saturday || weekdate == time.Sunday {
				date = date.AddDate(0, 0, 1)
			} else {
				date = date.AddDate(0, 0, 1)
				num--
				count++
			}
		}
	} else {

		for num < 0 {

			weekdate := date.Weekday()
			if weekdate == time.Saturday || weekdate == time.Sunday {
				date = date.AddDate(0, 0, -1)
			} else {
				date = date.AddDate(0, 0, -1)
				num++
				count++
			}
		}
	}
	fmt.Println(date)
}
