package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	// location, _ := time.LoadLocation("Asia/Kolkata")
	// currentTime := time.Now()
	// currentTime = currentTime.In(location)
	// fmt.Println(currentTime.Format("2006 Jan 02 Monday"))

	var str1 string
	var str2 string
	fmt.Println(`Enter date 1st in Format "YYYY-MM-DD" : `)
	fmt.Scanf("%v", &str1)

	startdate, err := time.Parse("2006-01-02", str1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(`Enter date 2st in Format "YYYY-MM-DD" : `)
	fmt.Scanf("%v", &str2)

	enddate, err := time.Parse("2006-01-02", str2)

	if err != nil {
		log.Fatal(err)
	}
	var count int

	for date := startdate; date.Before(enddate); date = date.AddDate(0, 0, 1) {

		day := date.Weekday()
		if day == time.Saturday || day == time.Sunday {
			count++
		}
	}
	fmt.Println(count)
}
